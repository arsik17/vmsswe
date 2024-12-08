package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Email       string             `json:"email"`
	FirstName   string             `json:"first_name"`
	LastName    string             `json:"last_name"`
	MiddleName  string             `json:"middle_name"`
	Password    string             `json:"password"`
	RoleID      primitive.ObjectID `json:"role_id"`
	GovermentID int                `json:"goverment_id"`
	PhoneNumber string             `json:"phone_number"`
	Address     string             `json:"address"`
}

type CreateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permissions int    `json:"permissions"`
}

type CreateRouteRequest struct {
	DriverID      primitive.ObjectID `json:"driver_id"`
	StartLocation string             `json:"start_location"`
	EndLocation   string             `json:"end_location"`
	Description   string             `json:"description"`
	TimeForRoute  int                `json:"time_for_route"`
	Distance      float32            `json:"distance"`
	Status        string             `json:"status"`
}

type CreateFuelingRequest struct {
	FuelerID   primitive.ObjectID `json:"fueler_id"`
	VehicleID  string             `json:"vehicle_id"`
	FuelAmount float32            `json:"fuel_amount"`
	TotalCost  float32            `json:"total_cost"`
	Year       int                `json:"year"`
	Month      int                `json:"month"`
	Day        int                `json:"day"`
	GasStation string             `json:"gas_station"`
}

type CreateMaintenanceRequest struct {
	MaintenancePersonID primitive.ObjectID `json:"maintenance_person_id"`
	VehicleID           string             `json:"vehicle_id"`
	Year                int                `json:"year"`
	Month               int                `json:"month"`
	Day                 int                `json:"day"`
	Parts               []string           `json:"parts"`
	ServiceType         string             `json:"service_type"`
	Proof               string             `json:"maintenance_proof"`
	Cost                float32            `json:"cost"`
}

type CreateVehicleRequest struct {
	VIN            string             `json:"_id"`
	ProductionYear int                `json:"production_year"`
	Mileage        int                `json:"mileage"`
	LicensePlate   string             `json:"license_plate"`
	DriverID       primitive.ObjectID `json:"driver_id"`
	Model          string             `json:"model"`
	CarMake        string             `json:"car_make"`
	ActivityStatus string             `json:"activity_status"`
	ExactLocation  string             `json:"exact_location"`
}

type DriverResponce struct {
	ID          primitive.ObjectID `json:"_id"`
	Email       string             `json:"email"`
	FirstName   string             `json:"first_name"`
	LastName    string             `json:"last_name"`
	MiddleName  string             `json:"middle_name"`
	Password    string             `json:"password"`
	RoleID      primitive.ObjectID `json:"role_id"`
	GovermentID int                `json:"goverment_id"`
	PhoneNumber string             `json:"phone_number"`
	Address     string             `json:"address"`
}

type UserResponce struct {
	ID          primitive.ObjectID `json:"_id"`
	Email       string             `json:"email"`
	FirstName   string             `json:"first_name"`
	LastName    string             `json:"last_name"`
	MiddleName  string             `json:"middle_name"`
	Password    string             `json:"password"`
	RoleID      primitive.ObjectID `json:"role_id"`
	GovermentID int                `json:"goverment_id"`
	PhoneNumber string             `json:"phone_number"`
}

type CreateAuctionRequest struct {
	VehicleID     string `json:"vehicle_id" bson:"vehicle_id"`
	Name          string `json:"name" bson:"name"`
	Description   string `json:"description" bson:"description"`
	StartingPrice int    `json:"starting_price" bson:"starting_price"`
	CurrentPrice  int    `json:"current_price" bson:"current_price"`
	Status        string `json:"status" bson:"status"`
}
type CreateReportRequest struct {
	TotalDistance int                `bson:"total_distance" json:"total_distance"`
	MoneySpent    int                `bson:"money_spent" json:"money_spent"`
	FuelUsage     float32            `bson:"fuel_usage" json:"fuel_usage"`
	DriverID      primitive.ObjectID `bson:"driver_id" json:"driver_id"`
	VehicleID     string             `bson:"vehicle_id" json:"vehicle_id"`
	RouteID       primitive.ObjectID `bson:"route_id" json:"route_id"`
}

type DeleteRequest struct {
	ID primitive.ObjectID `json:"_id"`
}

type DeleteRequestVehicle struct {
	ID string `json:"_id"`
}
