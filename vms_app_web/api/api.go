package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	model "github.com/arystanbek2002/swe/model"
	stat "github.com/arystanbek2002/swe/status"
	"github.com/arystanbek2002/swe/storage"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type APIServer struct {
	listenAddr string
	Store      storage.Storage
}

type APIError struct {
	Error string `json:"error"`
}

type Claims struct {
	ID          primitive.ObjectID `json:"id"`
	RoleID      primitive.ObjectID `json:"role_id"`
	Permissions int                `json:"permission"`
	jwt.RegisteredClaims
}

type void struct{}

var ctx = context.Background()

func NewAPIServer(addr string, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: addr,
		Store:      store,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Println(v)
		return fmt.Errorf(stat.EncodingError)
	}
	return nil
}

type APIFunc func(w http.ResponseWriter, r *http.Request, claims *Claims) error

func makeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r, nil); err != nil {
			//actually status cant be changed as it set but to invoke the func status must be used as argument
			WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
	}
}

func (api *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/createUser", ValidateJWT(api.handleCreateUser, map[int]void{1: {}}))
	router.HandleFunc("/createRole", ValidateJWT(api.handleCreateRole, map[int]void{1: {}}))
	router.HandleFunc("/createFueling", ValidateJWT(api.handleCreateFueling, map[int]void{1: {}}))
	router.HandleFunc("/createRoute", ValidateJWT(api.handleCreateRoute, map[int]void{1: {}}))
	router.HandleFunc("/createMaintenance", ValidateJWT(api.handleCreateMaintenance, map[int]void{1: {}}))
	router.HandleFunc("/createVehicle", ValidateJWT(api.handleCreateVehicle, map[int]void{1: {}}))
	router.HandleFunc("/createAuction", ValidateJWT(api.handleCreateAuction, map[int]void{1: {}}))
	router.HandleFunc("/createReport", ValidateJWT(api.handleCreateReport, map[int]void{1: {}}))

	router.HandleFunc("/loginUser", makeHTTPHandleFunc(api.handleLogin))
	router.HandleFunc("/loginAdmin", makeHTTPHandleFunc(api.handleAdminLogin))

	router.HandleFunc("/users", ValidateJWT(api.getUsers, map[int]void{1: {}}))
	router.HandleFunc("/roles", ValidateJWT(api.getRoles, map[int]void{1: {}}))
	router.HandleFunc("/fuelers", ValidateJWT(api.getFuelers, map[int]void{1: {}}))
	router.HandleFunc("/drivers", ValidateJWT(api.getDrivers, map[int]void{1: {}}))
	router.HandleFunc("/maintainers", ValidateJWT(api.getMaintainers, map[int]void{1: {}}))
	router.HandleFunc("/routes", ValidateJWT(api.getRoutes, map[int]void{1: {}}))
	router.HandleFunc("/vehicles", ValidateJWT(api.getVehicles, map[int]void{1: {}}))
	router.HandleFunc("/fuelings", ValidateJWT(api.getFuelings, map[int]void{1: {}}))
	router.HandleFunc("/maintenances", ValidateJWT(api.GetMaintenances, map[int]void{1: {}}))
	router.HandleFunc("/auctions", ValidateJWT(api.getAuctions, map[int]void{1: {}}))
	router.HandleFunc("/reports", ValidateJWT(api.getReports, map[int]void{1: {}}))

	router.HandleFunc("/getDriverRoutes", ValidateJWT(api.getRoutesByDriver, map[int]void{1: {}, 2: {}}))
	router.HandleFunc("/getDriverActiveRoutes", ValidateJWT(api.getActiveRoutesByDriver, map[int]void{1: {}, 2: {}}))
	router.HandleFunc("/cancelRoute", ValidateJWT(api.CancelRoute, map[int]void{1: {}, 2: {}}))
	router.HandleFunc("/completeRoute", ValidateJWT(api.CompleteRoute, map[int]void{1: {}, 2: {}}))

	router.HandleFunc("/deleteDriver", ValidateJWT(api.deleteDriver, map[int]void{1: {}}))
	router.HandleFunc("/deleteFueler", ValidateJWT(api.deleteFueler, map[int]void{1: {}}))
	router.HandleFunc("/deleteMaintainer", ValidateJWT(api.deleteMaintainer, map[int]void{1: {}}))
	router.HandleFunc("/deleteVehicle", ValidateJWT(api.deleteVehicle, map[int]void{1: {}}))
	router.HandleFunc("/deleteFueling", ValidateJWT(api.deleteFueling, map[int]void{1: {}}))
	router.HandleFunc("/deleteMaintenance", ValidateJWT(api.deleteMaintenance, map[int]void{1: {}}))
	router.HandleFunc("/deleteRoutes", ValidateJWT(api.deleteRoutes, map[int]void{1: {}}))
	router.HandleFunc("/deleteVehicle", ValidateJWT(api.deleteVehicle, map[int]void{1: {}}))
	router.HandleFunc("/deleteAuction", ValidateJWT(api.deleteAuction, map[int]void{1: {}}))
	router.HandleFunc("/deleteReport", ValidateJWT(api.deleteReport, map[int]void{1: {}}))

	router.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/",
		http.FileServer(http.Dir("styles"))))

	router.PathPrefix("/scripts/").Handler(http.StripPrefix("/scripts/",
		http.FileServer(http.Dir("scripts"))))

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",
		http.FileServer(http.Dir("assets"))))

	router.PathPrefix("/").Handler(http.StripPrefix("/",
		http.FileServer(http.Dir("static"))))

	if err := http.ListenAndServe(api.listenAddr, router); err != nil {
		log.Printf("error starting server %s", err.Error())
	}
	fmt.Println("server start running")
}

func validateUserRequest(req *model.UserRequest) (int, error) {
	if len(req.Email) < 1 {
		return http.StatusBadRequest, fmt.Errorf("wrong credentials")
	}
	if len(req.Password) < 1 {
		return http.StatusBadRequest, fmt.Errorf("wrong credentials")
	}
	return http.StatusOK, nil
}

func validateCreateUserRequest(req *model.CreateUserRequest) (int, error) {
	if len(req.Email) < 6 {
		return http.StatusBadRequest, fmt.Errorf("wrong credentials")
	}
	if len(req.Password) < 5 {
		return http.StatusBadRequest, fmt.Errorf("wrong credentials")
	}
	if len(req.FirstName) == 0 {
		return http.StatusBadRequest, fmt.Errorf("name must not be nil")
	}
	if len(req.LastName) == 0 {
		return http.StatusBadRequest, fmt.Errorf("surname must not be nil")
	}

	return http.StatusOK, nil
}

func validateCreateRoleRequest(req *model.CreateRoleRequest) (int, error) {
	if req.Description == "" {
		return http.StatusBadRequest, fmt.Errorf("description must not be nil")
	}
	if req.Name == "" {
		return http.StatusBadRequest, fmt.Errorf("name must not be nil")
	}
	if req.Permissions == 0 {
		return http.StatusBadRequest, fmt.Errorf("permissions must not be nil")
	}
	return http.StatusOK, nil
}

func validateCreateRouteRequest(req *model.CreateRouteRequest) (int, error) {
	if req.DriverID == primitive.NilObjectID {
		return http.StatusBadRequest, fmt.Errorf("DriverID must not be nil")
	}
	if req.Description == "" {
		return http.StatusBadRequest, fmt.Errorf("description must not be nil")
	}
	if req.EndLocation == "" {
		return http.StatusBadRequest, fmt.Errorf("end location must not be nil")
	}
	if req.StartLocation == "" {
		return http.StatusBadRequest, fmt.Errorf("start location must not be nil")
	}
	if req.Status == "" {
		return http.StatusBadRequest, fmt.Errorf("status must not be nil")
	}
	if req.TimeForRoute == 0 {
		return http.StatusBadRequest, fmt.Errorf("time for route must not be nil")
	}
	if req.Distance <= 0 {
		return http.StatusBadRequest, fmt.Errorf("distance must not be nil or negative")
	}
	return http.StatusOK, nil
}

func validateCreateFuelingRequest(req *model.CreateFuelingRequest) (int, error) {
	if req.FuelerID == primitive.NilObjectID {
		return http.StatusBadRequest, fmt.Errorf("FuelerID must not be nil")
	}
	if req.VehicleID == "" {
		return http.StatusBadRequest, fmt.Errorf("VehicleID must not be nil")
	}
	if req.FuelAmount <= 0 {
		return http.StatusBadRequest, fmt.Errorf("FuelAmount location must not be nil")
	}
	if req.TotalCost <= 0 {
		return http.StatusBadRequest, fmt.Errorf("TotalCost location must not be nil")
	}
	if req.GasStation == "" {
		return http.StatusBadRequest, fmt.Errorf("GasStation must not be nil")
	}
	if req.Year == 0 {
		return http.StatusBadRequest, fmt.Errorf("year must not be nil")
	}
	if req.Month == 0 {
		return http.StatusBadRequest, fmt.Errorf("month must not be nil")
	}
	if req.Day == 0 {
		return http.StatusBadRequest, fmt.Errorf("day must not be nil")
	}
	return http.StatusOK, nil
}

func validateCreateMaintenanceRequest(req *model.CreateMaintenanceRequest) (int, error) {
	if req.MaintenancePersonID == primitive.NilObjectID {
		return http.StatusBadRequest, fmt.Errorf("FuelerID must not be nil")
	}
	if req.VehicleID == "" {
		return http.StatusBadRequest, fmt.Errorf("VehicleID must not be nil")
	}
	if len(req.Parts) == 0 {
		return http.StatusBadRequest, fmt.Errorf("FuelAmount location must not be nil")
	}
	if req.ServiceType == "" {
		return http.StatusBadRequest, fmt.Errorf("TotalCost location must not be nil")
	}
	if req.Cost < 0 {
		return http.StatusBadRequest, fmt.Errorf("GasStation must not be nil")
	}
	if req.Year == 0 {
		return http.StatusBadRequest, fmt.Errorf("year must not be nil")
	}
	if req.Month < 1 || req.Month > 12 {
		return http.StatusBadRequest, fmt.Errorf("month must not be nil")
	}
	if req.Day < 1 || req.Day > 31 {
		return http.StatusBadRequest, fmt.Errorf("day must not be nil")
	}
	return http.StatusOK, nil
}

func validateCreateVehicleRequest(req *model.CreateVehicleRequest) (int, error) {
	if len(req.VIN) < 3 {
		return http.StatusBadRequest, fmt.Errorf("FuelerID must not be nil")
	}
	if req.Mileage < 0 {
		return http.StatusBadRequest, fmt.Errorf("FuelAmount location must not be nil")
	}
	if req.LicensePlate == "" {
		return http.StatusBadRequest, fmt.Errorf("TotalCost location must not be nil %v 3", req.LicensePlate)
	}
	if req.Model == "" {
		return http.StatusBadRequest, fmt.Errorf("year must not be nil")
	}
	if req.CarMake == "" {
		return http.StatusBadRequest, fmt.Errorf("month must not be nil")
	}
	if req.ActivityStatus == "" {
		return http.StatusBadRequest, fmt.Errorf("day must not be nil")
	}
	return http.StatusOK, nil
}

func (api *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	var role *model.Role
	request := new(model.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	status, err := validateCreateUserRequest(request)
	if err != nil {
		log.Println(2)
		return WriteJSON(w, status, APIError{Error: err.Error()})
	}
	user := model.NewUser(request.Email, request.FirstName, request.LastName, request.MiddleName,
		request.Password, request.RoleID, request.GovermentID, request.PhoneNumber)

	if request.RoleID == primitive.NilObjectID {
		if err := api.Store.CreateUser(ctx, user); err != nil {
			return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
		return WriteJSON(w, http.StatusOK, "success")
	}
	role, err = api.Store.GetRole(ctx, request.RoleID)
	if err != nil {
		if err.Error() == stat.NoRole {
			return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	if err := api.Store.CreateUser(ctx, user); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	if role.Name == "driver" {
		if request.Address != "" {
			driver := model.NewDriver(user.ID, request.Address)
			if err := api.Store.AddDriver(context.Background(), driver); err != nil {
				return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
			}
		} else {
			return WriteJSON(w, http.StatusInternalServerError, APIError{Error: "address must not be nil while creating driver"})
		}
		return WriteJSON(w, http.StatusOK, "success")
	}
	if role.Name == "fueling person" {
		fueler := model.NewFuelingPerson(user.ID)
		if err := api.Store.AddFuelingPerson(context.Background(), fueler); err != nil {
			return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
		return WriteJSON(w, http.StatusOK, "success")
	}
	if role.Name == "maintenance person" {
		maintainer := model.NewMaintenancePerson(user.ID)
		if err := api.Store.AddMaintenancePerson(context.Background(), maintainer); err != nil {
			return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
		return WriteJSON(w, http.StatusOK, "success")
	}
	return WriteJSON(w, http.StatusInternalServerError, APIError{Error: "role has not yet been implemented"})
}

func (api *APIServer) handleCreateRole(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	request := new(model.CreateRoleRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	status, err := validateCreateRoleRequest(request)
	if err != nil {
		return WriteJSON(w, status, APIError{Error: err.Error()})
	}
	role := model.NewRole(request.Name, request.Description, request.Permissions)

	if _, err := api.Store.CreateRole(ctx, role); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) handleCreateRoute(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	request := new(model.CreateRouteRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	status, err := validateCreateRouteRequest(request)
	if err != nil {
		return WriteJSON(w, status, request)
	}
	route := model.NewRoute(request.DriverID, request.StartLocation, request.EndLocation, request.Description,
		request.TimeForRoute, request.Distance, request.Status)
	if err := api.Store.CancelRoute(ctx, request.DriverID); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	if err := api.Store.CreateRoute(ctx, route); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) handleCreateFueling(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	id, err := primitive.ObjectIDFromHex(r.FormValue("fueler_id"))
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	fuel, err := strconv.ParseFloat(r.FormValue("fuel_amount"), 32)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	cost, err := strconv.ParseFloat(r.FormValue("total_cost"), 32)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	month, err := strconv.Atoi(r.FormValue("month"))
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	day, err := strconv.Atoi(r.FormValue("day"))
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	request := &model.CreateFuelingRequest{
		FuelerID:   id,
		VehicleID:  r.FormValue("vehicle_id"),
		FuelAmount: float32(fuel),
		TotalCost:  float32(cost),
		Year:       year,
		Month:      month,
		Day:        day,
		GasStation: r.FormValue("gas_station"),
	}
	status, err := validateCreateFuelingRequest(request)
	if err != nil {
		return WriteJSON(w, status, APIError{Error: err.Error() + " 321"})
	}

	r.ParseMultipartForm(1 << 20)
	file, fileHeader, err := r.FormFile("fueling_proof")
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	defer file.Close()

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./fuelings", os.ModePerm)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	filename := fmt.Sprintf("./fuelings/%s%s", uuid.New(), filepath.Ext(fileHeader.Filename))
	// Create a new file in the uploads directory
	dst, err := os.Create(filename)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	fueling := model.NewFueling(request.FuelerID, request.VehicleID, request.FuelAmount, request.TotalCost,
		time.Date(request.Year, time.Month(request.Month), request.Day, 0, 0, 0, 0, time.UTC), request.GasStation, filename)

	if err := api.Store.CreateFueling(ctx, fueling); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) handleCreateMaintenance(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	id, err := primitive.ObjectIDFromHex(r.FormValue("maintenance_person_id"))
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	cost, err := strconv.ParseFloat(r.FormValue("cost"), 32)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	month, err := strconv.Atoi(r.FormValue("month"))
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	day, err := strconv.Atoi(r.FormValue("day"))
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	request := &model.CreateMaintenanceRequest{
		MaintenancePersonID: id,
		VehicleID:           r.FormValue("vehicle_id"),
		Parts:               strings.Split(r.FormValue("parts"), ","),
		ServiceType:         r.FormValue("service_type"),
		Year:                year,
		Month:               month,
		Day:                 day,
		Cost:                float32(cost),
	}
	status, err := validateCreateMaintenanceRequest(request)
	if err != nil {
		return WriteJSON(w, status, APIError{Error: err.Error()})
	}

	r.ParseMultipartForm(1 << 20)
	file, fileHeader, err := r.FormFile("maintenance_proof")
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	defer file.Close()

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./maintains", os.ModePerm)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	filename := fmt.Sprintf("./maintains/%s%s", uuid.New(), filepath.Ext(fileHeader.Filename))
	// Create a new file in the uploads directory
	dst, err := os.Create(filename)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	maintaining := model.NewMaintenance(request.MaintenancePersonID, request.VehicleID, request.Parts, request.ServiceType,
		time.Date(request.Year, time.Month(request.Month), request.Day, 0, 0, 0, 0, time.UTC), filename, request.Cost)

	if err := api.Store.CreateMaintenance(ctx, maintaining); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) handleCreateAuction(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	request := new(model.CreateAuctionRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	auction := model.NewAuction(request.VehicleID, request.Name, request.Description, request.StartingPrice,
		request.CurrentPrice, request.Status)
	if err := api.Store.CreateAuction(ctx, auction); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) handleCreateReport(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	request := new(model.CreateReportRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	report := model.NewReport(request.TotalDistance, request.MoneySpent, request.FuelUsage, request.DriverID,
		request.VehicleID, request.RouteID)
	if err := api.Store.CreateReport(ctx, report); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) handleLogin(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

	request := new(model.UserRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusOK, APIError{Error: "1"})
	}
	status, err := validateUserRequest(request)
	if err != nil {
		return WriteJSON(w, status, APIError{Error: "2"})
	}

	user, err := api.Store.LoginUser(ctx, request.Email, request.Password)
	if err != nil {
		if err.Error() == stat.WrongCredentials {
			return WriteJSON(w, http.StatusBadRequest, APIError{Error: "3"})
		}
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	role, err := api.Store.GetRole(ctx, user.RoleID)
	if err != nil {
		if err.Error() == stat.NoRole {
			return WriteJSON(w, http.StatusBadRequest, APIError{Error: "5"})
		}
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: "6"})
	}
	responce := model.DriverResponce{}
	if role.Name == "driver" {
		driver, err := api.Store.GetDriver(ctx, user.ID)
		if err != nil {
			return WriteJSON(w, status, APIError{Error: err.Error()})
		}
		responce.Address = driver.Adress
	}

	responce.Email = user.Email
	responce.FirstName = user.FirstName
	responce.GovermentID = user.GovermentID
	responce.ID = user.ID
	responce.LastName = user.LastName
	responce.MiddleName = user.MiddleName
	responce.PhoneNumber = user.PhoneNumber
	responce.RoleID = user.RoleID

	jwToken, expirationTime, err := generateJWT(user, role, time.Now())
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: "7"})
	}

	cookie := http.Cookie{
		Name:     "x-jwt",
		Value:    jwToken,
		Path:     "/",
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, &cookie)
	return WriteJSON(w, http.StatusOK, responce)
}

func (api *APIServer) handleAdminLogin(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

	request := new(model.UserRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	status, err := validateUserRequest(request)
	if err != nil {
		return WriteJSON(w, status, APIError{Error: err.Error()})
	}

	user, err := api.Store.LoginUser(ctx, request.Email, request.Password)
	if err != nil {
		if err.Error() == stat.WrongCredentials {
			return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	role, err := api.Store.GetRole(ctx, user.RoleID)
	if err != nil {
		if err.Error() == stat.NoRole {
			return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	if role.Name != "admin" {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: "not admin account"})
	}

	jwToken, expirationTime, err := generateJWT(user, role, time.Now())
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	cookie := http.Cookie{
		Name:     "x-jwt",
		Value:    jwToken,
		Path:     "/",
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, &cookie)
	return WriteJSON(w, http.StatusOK, APIError{Error: "success"})
}

func (api *APIServer) handleCreateVehicle(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	request := new(model.CreateVehicleRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	status, err := validateCreateVehicleRequest(request)
	if err != nil {
		return WriteJSON(w, status, APIError{Error: err.Error()})
	}
	vehicle := model.NewVehicle(request.VIN, request.ProductionYear,
		request.Mileage, request.LicensePlate,
		request.DriverID, request.Model, request.CarMake, request.ActivityStatus, request.ExactLocation)

	if err := api.Store.CreateVehicle(ctx, vehicle); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func generateJWT(user *model.User, role *model.Role, t time.Time) (string, time.Time, error) {
	expirationTime := t.Add(12 * time.Hour)

	claims := &Claims{
		ID:          user.ID,
		RoleID:      user.RoleID,
		Permissions: role.Permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", time.Now(), err
	}
	return tokenString, expirationTime, nil
}

func VerifyJWT(jwtString string) (int, *Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return http.StatusUnauthorized, nil, err //fmt.Errorf(errStatus.BadToken)
	}
	if !tkn.Valid {
		return http.StatusUnauthorized, nil, fmt.Errorf(stat.InvalidToken)
	}
	return http.StatusOK, claims, nil
}

func ValidateJWT(h APIFunc, permissions map[int]void) http.HandlerFunc {
	return http.HandlerFunc(makeHTTPHandleFunc(func(w http.ResponseWriter, r *http.Request, v *Claims) error {
		c, err := r.Cookie("x-jwt")
		if err != nil {
			return WriteJSON(w, http.StatusUnauthorized, APIError{Error: stat.NoCookie})
		}

		tknStr := c.Value
		status, claims, err := VerifyJWT(tknStr)

		_, ok := permissions[claims.Permissions]
		if !ok {
			return WriteJSON(w, status, APIError{Error: "No permission"})
		}

		if err != nil {
			return WriteJSON(w, status, APIError{Error: err.Error()})
		}
		h(w, r, claims)
		return nil
	}))
}

func (api *APIServer) getUsers(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	users, err := api.Store.GetUsers(ctx)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, users)
}

func (api *APIServer) getRoles(w http.ResponseWriter, r *http.Request, _ *Claims) error {

	roles, err := api.Store.GetRoles(ctx)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, roles)
}

func (api *APIServer) getDrivers(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	drivers, err := api.Store.GetDrivers(ctx)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	responces := []*model.DriverResponce{}
	for _, driver := range drivers {
		user, err := api.Store.GetUser(ctx, driver.ID)
		if err != nil {
			return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
		responce := &model.DriverResponce{
			ID:          user.ID,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			MiddleName:  user.MiddleName,
			Password:    user.Password,
			RoleID:      user.RoleID,
			GovermentID: user.GovermentID,
			PhoneNumber: user.PhoneNumber,
			Address:     driver.Adress,
		}
		responces = append(responces, responce)
	}
	return WriteJSON(w, http.StatusOK, responces)
}

func (api *APIServer) getFuelers(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	fuelers, err := api.Store.GetFuelers(ctx)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	responces := []*model.UserResponce{}
	for _, fueler := range fuelers {
		user, err := api.Store.GetUser(ctx, fueler.ID)
		if err != nil {
			return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
		responce := &model.UserResponce{
			ID:          user.ID,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			MiddleName:  user.MiddleName,
			Password:    user.Password,
			RoleID:      user.RoleID,
			GovermentID: user.GovermentID,
			PhoneNumber: user.PhoneNumber,
		}
		responces = append(responces, responce)
	}
	return WriteJSON(w, http.StatusOK, responces)
}

func (api *APIServer) getMaintainers(w http.ResponseWriter, r *http.Request, _ *Claims) error {
	maintainers, err := api.Store.GetMaintainers(ctx)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	responces := []*model.UserResponce{}
	for _, maintainer := range maintainers {
		user, err := api.Store.GetUser(ctx, maintainer.ID)
		if err != nil {
			return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
		responce := &model.UserResponce{
			ID:          user.ID,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			MiddleName:  user.MiddleName,
			Password:    user.Password,
			RoleID:      user.RoleID,
			GovermentID: user.GovermentID,
			PhoneNumber: user.PhoneNumber,
		}
		responces = append(responces, responce)
	}
	return WriteJSON(w, http.StatusOK, responces)
}

func (api *APIServer) getFuelings(w http.ResponseWriter, r *http.Request, _ *Claims) error {

	fuelings, err := api.Store.GetFuelings(ctx)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, fuelings)
}

func (api *APIServer) getRoutes(w http.ResponseWriter, r *http.Request, _ *Claims) error {

	routes, err := api.Store.GetRoutes(ctx)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, routes)
}

func (api *APIServer) GetMaintenances(w http.ResponseWriter, r *http.Request, _ *Claims) error {

	maintenances, err := api.Store.GetMaintenances(ctx)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, maintenances)
}

func (api *APIServer) getVehicles(w http.ResponseWriter, r *http.Request, _ *Claims) error {

	vehicles, err := api.Store.GetVehicles(ctx)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, vehicles)
}

func (api *APIServer) getAuctions(w http.ResponseWriter, r *http.Request, _ *Claims) error {

	auctions, err := api.Store.GetAuctions(ctx)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, auctions)
}

func (api *APIServer) getReports(w http.ResponseWriter, r *http.Request, _ *Claims) error {

	reports, err := api.Store.GetReports(ctx)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, reports)
}

func (api *APIServer) getRoutesByDriver(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	oid := claims.ID

	driver, err := api.Store.GetDriver(ctx, oid)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	routes, err := api.Store.GetRoutesByDriver(ctx, driver.ID)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, routes)
}

func (api *APIServer) getActiveRoutesByDriver(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	oid := claims.ID

	driver, err := api.Store.GetDriver(ctx, oid)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	route, err := api.Store.GetActiveRoutesByDriver(ctx, driver.ID)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, route)
}

func (api *APIServer) CancelRoute(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	oid := claims.ID

	driver, err := api.Store.GetDriver(ctx, oid)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	err = api.Store.CancelRoute(ctx, driver.ID)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) CompleteRoute(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	oid := claims.ID

	driver, err := api.Store.GetDriver(ctx, oid)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	err = api.Store.CompleteRoute(ctx, driver.ID)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) deleteDriver(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	request := new(model.DeleteRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	err := api.Store.DeleteDriver(ctx, request.ID)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) deleteFueler(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	request := new(model.DeleteRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	err := api.Store.DeleteFueler(ctx, request.ID)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) deleteMaintainer(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	request := new(model.DeleteRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	err := api.Store.DeleteMaintainer(ctx, request.ID)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) deleteVehicle(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	request := new(model.DeleteRequestVehicle)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	err := api.Store.DeleteVehicle(ctx, request.ID)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) deleteFueling(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	request := new(model.DeleteRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	err := api.Store.DeleteFueling(ctx, request.ID)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) deleteMaintenance(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	request := new(model.DeleteRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	err := api.Store.DeleteMaintenance(ctx, request.ID)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) deleteRoutes(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	request := new(model.DeleteRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	err := api.Store.DeleteRoutes(ctx, request.ID)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) deleteAuction(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	request := new(model.DeleteRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	err := api.Store.DeleteAuction(ctx, request.ID)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}

func (api *APIServer) deleteReport(w http.ResponseWriter, r *http.Request, claims *Claims) error {
	request := new(model.DeleteRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	err := api.Store.DeleteReport(ctx, request.ID)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "success")
}
