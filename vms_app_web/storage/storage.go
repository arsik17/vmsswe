package storage

import (
	"context"
	"fmt"

	"github.com/arystanbek2002/swe/model"
	stat "github.com/arystanbek2002/swe/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface {
	CreateUser(context.Context, *model.User) error
	AddAdmin(context.Context, *model.Admin) error
	AddDriver(context.Context, *model.Driver) error
	CreateRole(context.Context, *model.Role) (interface{}, error)
	CreateRoute(context.Context, *model.Route) error
	AddFuelingPerson(context.Context, *model.FuelingPerson) error
	CreateFueling(context.Context, *model.Fueling) error
	AddMaintenancePerson(context.Context, *model.MaintenancePerson) error
	CreateMaintenance(context.Context, *model.Maintenance) error
	CreateVehicle(context.Context, *model.Vehicle) error
	CreateAuction(context.Context, *model.Auction) error
	CreateReport(context.Context, *model.Report) error

	GetUsers(context.Context) ([]*model.User, error)
	GetAdmins(context.Context) ([]*model.Admin, error)
	GetDrivers(context.Context) ([]*model.Driver, error)
	GetRoles(context.Context) ([]*model.Role, error)
	GetRoutes(context.Context) ([]*model.Route, error)
	GetFuelers(context.Context) ([]*model.FuelingPerson, error)
	GetFuelings(context.Context) ([]*model.Fueling, error)
	GetMaintainers(context.Context) ([]*model.MaintenancePerson, error)
	GetMaintenances(context.Context) ([]*model.Maintenance, error)
	GetVehicles(context.Context) ([]*model.Vehicle, error)
	GetAuctions(context.Context) ([]*model.Auction, error)
	GetReports(context.Context) ([]*model.Report, error)

	GetUser(context.Context, primitive.ObjectID) (*model.User, error)
	GetAdmin(context.Context, primitive.ObjectID) (*model.Admin, error)
	GetDriver(context.Context, primitive.ObjectID) (*model.Driver, error)
	GetRole(context.Context, primitive.ObjectID) (*model.Role, error)
	GetRoute(context.Context, primitive.ObjectID) (*model.Route, error)
	GetFueler(context.Context, primitive.ObjectID) (*model.FuelingPerson, error)
	GetFueling(context.Context, primitive.ObjectID) (*model.Fueling, error)
	GetMaintainer(context.Context, primitive.ObjectID) (*model.MaintenancePerson, error)
	GetMaintenance(context.Context, primitive.ObjectID) (*model.Maintenance, error)
	GetVehicle(context.Context, primitive.ObjectID) (*model.Vehicle, error)
	GetAuction(context.Context, primitive.ObjectID) (*model.Auction, error)
	GetReport(context.Context, primitive.ObjectID) (*model.Report, error)

	LoginUser(context.Context, string, string) (*model.User, error)
	GetRoutesByDriver(context.Context, primitive.ObjectID) ([]*model.Route, error)
	GetActiveRoutesByDriver(context.Context, primitive.ObjectID) (*model.Route, error)
	CancelRoute(context.Context, primitive.ObjectID) error
	CompleteRoute(context.Context, primitive.ObjectID) error

	DeleteUser(context.Context, primitive.ObjectID) error
	DeleteDriver(context.Context, primitive.ObjectID) error
	DeleteFueler(context.Context, primitive.ObjectID) error
	DeleteMaintainer(context.Context, primitive.ObjectID) error
	DeleteVehicle(context.Context, string) error
	DeleteFueling(context.Context, primitive.ObjectID) error
	DeleteMaintenance(context.Context, primitive.ObjectID) error
	DeleteRoutes(context.Context, primitive.ObjectID) error
	DeleteAuction(context.Context, primitive.ObjectID) error
	DeleteReport(context.Context, primitive.ObjectID) error
}

type MongoStore struct {
	users        *mongo.Collection
	admins       *mongo.Collection
	drivers      *mongo.Collection
	roles        *mongo.Collection
	routes       *mongo.Collection
	fuelers      *mongo.Collection
	fuelings     *mongo.Collection
	mainPersons  *mongo.Collection
	maintenances *mongo.Collection
	vehicles     *mongo.Collection
	auctions     *mongo.Collection
	reports      *mongo.Collection
}

func NewMongoStore(ctx context.Context, client *mongo.Client) (*MongoStore, error) {
	client.Database("swe").Drop(ctx)
	str := &MongoStore{
		users:        client.Database("swe").Collection("users"),
		admins:       client.Database("swe").Collection("admins"),
		drivers:      client.Database("swe").Collection("drivers"),
		roles:        client.Database("swe").Collection("roles"),
		routes:       client.Database("swe").Collection("routes"),
		fuelers:      client.Database("swe").Collection("fuelers"),
		fuelings:     client.Database("swe").Collection("fueling"),
		mainPersons:  client.Database("swe").Collection("mainPersons"),
		maintenances: client.Database("swe").Collection("maintenances"),
		vehicles:     client.Database("swe").Collection("vehicles"),
		auctions:     client.Database("swe").Collection("auctions"),
		reports:      client.Database("swe").Collection("reports"),
	}
	role := model.NewRole("admin", "asd", 1)
	_, err := str.CreateRole(ctx, role)
	if err != nil {
		return nil, err
	}
	user := model.NewUser("admin@admin.com", "a", "a", "", "zxc", role.ID, 123121231, "sdsdfsdfs")
	err = str.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	_, err = str.CreateRole(ctx, model.NewRole("driver", "asd", 2))
	if err != nil {
		return nil, err
	}
	_, err = str.CreateRole(ctx, model.NewRole("fueling person", "asd", 3))
	if err != nil {
		return nil, err
	}
	_, err = str.CreateRole(ctx, model.NewRole("maintenance person", "asd", 4))
	if err != nil {
		return nil, err
	}
	err = str.AddAdmin(ctx, model.NewAdmin(user.ID))
	if err != nil {
		return nil, err
	}
	return str, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (str *MongoStore) CreateUser(ctx context.Context, user *model.User) error {
	if user.RoleID == primitive.NilObjectID {
		filter := bson.D{{"_id", user.RoleID}}
		var space model.Role
		err := str.roles.FindOne(ctx, filter).Decode(&space)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return fmt.Errorf("role does not exists")
			}
			return fmt.Errorf("db internal err: %s", err)
		}
	}
	_, err := str.users.InsertOne(ctx, user)
	return err
}

func (str *MongoStore) AddAdmin(ctx context.Context, admin *model.Admin) error {
	filter := bson.D{{"_id", admin.ID}}
	var space model.User
	err := str.users.FindOne(ctx, filter).Decode(&space)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}
	_, err = str.admins.InsertOne(ctx, admin)
	return err
}

func (str *MongoStore) AddDriver(ctx context.Context, driver *model.Driver) error {
	filter := bson.D{{"_id", driver.ID}}
	var space model.User
	err := str.users.FindOne(ctx, filter).Decode(&space)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}
	_, err = str.drivers.InsertOne(ctx, driver)
	return err
}

func (str *MongoStore) CreateRole(ctx context.Context, role *model.Role) (interface{}, error) {
	oid, err := str.roles.InsertOne(ctx, role)
	return oid.InsertedID, err
}

func (str *MongoStore) CreateRoute(ctx context.Context, route *model.Route) error {
	filter := bson.D{{"_id", route.DriverID}}
	var space model.Driver
	err := str.drivers.FindOne(ctx, filter).Decode(&space)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("driver does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}
	_, err = str.routes.InsertOne(ctx, route)
	return err
}

func (str *MongoStore) AddFuelingPerson(ctx context.Context, fueler *model.FuelingPerson) error {
	filter := bson.D{{"_id", fueler.ID}}
	var space model.User
	err := str.users.FindOne(ctx, filter).Decode(&space)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}
	_, err = str.fuelers.InsertOne(ctx, fueler)
	return err
}

func (str *MongoStore) CreateFueling(ctx context.Context, fueling *model.Fueling) error {
	filter := bson.D{{"_id", fueling.FuelerID}}
	var space model.FuelingPerson
	err := str.fuelers.FindOne(ctx, filter).Decode(&space)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("fueler does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}

	filter = bson.D{{"_id", fueling.VehicleID}}
	var space2 model.Vehicle
	err = str.vehicles.FindOne(ctx, filter).Decode(&space2)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("vehicle does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}

	_, err = str.fuelings.InsertOne(ctx, fueling)
	return err
}

func (str *MongoStore) AddMaintenancePerson(ctx context.Context, mainPerson *model.MaintenancePerson) error {
	filter := bson.D{{"_id", mainPerson.ID}}
	var space model.User
	err := str.users.FindOne(ctx, filter).Decode(&space)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}
	_, err = str.mainPersons.InsertOne(ctx, mainPerson)
	return err
}

func (str *MongoStore) CreateMaintenance(ctx context.Context, maintenance *model.Maintenance) error {
	filter := bson.D{{"_id", maintenance.MaintenancePersonID}}
	var space model.MaintenancePerson
	err := str.mainPersons.FindOne(ctx, filter).Decode(&space)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("maintenance person does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}

	filter = bson.D{{"_id", maintenance.VehicleID}}
	var space2 model.Vehicle
	err = str.vehicles.FindOne(ctx, filter).Decode(&space2)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("vehicle does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}

	_, err = str.maintenances.InsertOne(ctx, maintenance)
	return err
}

func (str *MongoStore) CreateVehicle(ctx context.Context, vehicle *model.Vehicle) error {
	if vehicle.VIN == "" {
		return fmt.Errorf("VIN must not be nil")
	}

	if vehicle.DriverID != primitive.NilObjectID {
		filter := bson.D{{"_id", vehicle.DriverID}}
		var space model.Driver
		err := str.drivers.FindOne(ctx, filter).Decode(&space)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return fmt.Errorf("driver does not exists")
			}
			return fmt.Errorf("db internal err: %s", err)
		}
	}
	_, err := str.vehicles.InsertOne(ctx, vehicle)
	return err
}

func (str *MongoStore) CreateAuction(ctx context.Context, auction *model.Auction) error {
	filter := bson.D{{"_id", auction.VehicleID}}
	var space model.Vehicle
	err := str.vehicles.FindOne(ctx, filter).Decode(&space)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("vehicle does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}
	_, err = str.auctions.InsertOne(ctx, auction)
	return err
}

func (str *MongoStore) CreateReport(ctx context.Context, report *model.Report) error {
	filter := bson.D{{"_id", report.DriverID}}
	var space model.Driver
	err := str.drivers.FindOne(ctx, filter).Decode(&space)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("driver does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}

	filter = bson.D{{"_id", report.VehicleID}}
	var space2 model.Vehicle
	err = str.vehicles.FindOne(ctx, filter).Decode(&space2)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("vehicle does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}

	filter = bson.D{{"_id", report.RouteID}}
	var space3 model.Route
	err = str.routes.FindOne(ctx, filter).Decode(&space3)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("route does not exists")
		}
		return fmt.Errorf("db internal err: %s", err)
	}

	_, err = str.reports.InsertOne(ctx, report)
	return err
}
func get(col *mongo.Collection, result any, ctx context.Context) (any, error) {
	res, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	err = res.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (str *MongoStore) GetUsers(ctx context.Context) ([]*model.User, error) {
	var result []*model.User
	res, err := get(str.users, result, ctx)
	return res.([]*model.User), err
}

func (str *MongoStore) GetAdmins(ctx context.Context) ([]*model.Admin, error) {
	var result []*model.Admin
	res, err := get(str.admins, result, ctx)
	return res.([]*model.Admin), err
}

func (str *MongoStore) GetDrivers(ctx context.Context) ([]*model.Driver, error) {
	var result []*model.Driver
	res, err := get(str.drivers, result, ctx)
	return res.([]*model.Driver), err
}

func (str *MongoStore) GetRoles(ctx context.Context) ([]*model.Role, error) {
	var result []*model.Role
	res, err := get(str.roles, result, ctx)
	return res.([]*model.Role), err
}

func (str *MongoStore) GetRoutes(ctx context.Context) ([]*model.Route, error) {
	var result []*model.Route
	res, err := get(str.routes, result, ctx)
	return res.([]*model.Route), err
}

func (str *MongoStore) GetFuelers(ctx context.Context) ([]*model.FuelingPerson, error) {
	var result []*model.FuelingPerson
	res, err := get(str.fuelers, result, ctx)
	return res.([]*model.FuelingPerson), err
}

func (str *MongoStore) GetFuelings(ctx context.Context) ([]*model.Fueling, error) {
	var result []*model.Fueling
	res, err := get(str.fuelings, result, ctx)
	return res.([]*model.Fueling), err
}

func (str *MongoStore) GetMaintainers(ctx context.Context) ([]*model.MaintenancePerson, error) {
	var result []*model.MaintenancePerson
	res, err := get(str.mainPersons, result, ctx)
	return res.([]*model.MaintenancePerson), err
}

func (str *MongoStore) GetMaintenances(ctx context.Context) ([]*model.Maintenance, error) {
	var result []*model.Maintenance
	res, err := get(str.maintenances, result, ctx)
	return res.([]*model.Maintenance), err
}

func (str *MongoStore) GetVehicles(ctx context.Context) ([]*model.Vehicle, error) {
	var result []*model.Vehicle
	res, err := get(str.vehicles, result, ctx)
	return res.([]*model.Vehicle), err
}

func (str *MongoStore) GetAuctions(ctx context.Context) ([]*model.Auction, error) {
	var result []*model.Auction
	res, err := get(str.auctions, result, ctx)
	return res.([]*model.Auction), err
}

func (str *MongoStore) GetReports(ctx context.Context) ([]*model.Report, error) {
	var result []*model.Report
	res, err := get(str.reports, result, ctx)
	return res.([]*model.Report), err
}

func (str *MongoStore) LoginUser(ctx context.Context, email string, password string) (*model.User, error) {

	filter := bson.D{{"email", email}}
	var user model.User
	err := str.users.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("stat.WrongCredentials")
		}
		return nil, err
	}
	if password != user.Password {
		return nil, err
	}
	return &user, err
}

func getOne(col *mongo.Collection, result any, ctx context.Context, bsonTag string, id primitive.ObjectID) (any, error) {
	filter := bson.D{{bsonTag, id}}
	err := col.FindOne(ctx, filter).Decode(result)
	return result, err
}

func (str *MongoStore) GetUser(ctx context.Context, id primitive.ObjectID) (*model.User, error) {
	filter := bson.D{{"_id", id}}
	var user model.User
	err := str.users.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoUser)
		}
		return nil, err
	}
	return &user, err
}

func (str *MongoStore) GetAdmin(ctx context.Context, id primitive.ObjectID) (*model.Admin, error) {
	filter := bson.D{{"_id", id}}
	var admin model.Admin
	err := str.admins.FindOne(ctx, filter).Decode(&admin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoAdmin)
		}
		return nil, err
	}
	return &admin, err
}

func (str *MongoStore) GetDriver(ctx context.Context, id primitive.ObjectID) (*model.Driver, error) {
	filter := bson.D{{"_id", id}}
	var driver model.Driver
	err := str.drivers.FindOne(ctx, filter).Decode(&driver)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoDriver)
		}
		return nil, err
	}
	return &driver, err
}

func (str *MongoStore) GetRole(ctx context.Context, id primitive.ObjectID) (*model.Role, error) {
	filter := bson.D{{"_id", id}}
	var role model.Role
	err := str.roles.FindOne(ctx, filter).Decode(&role)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoRole)
		}
		return nil, err
	}
	return &role, err
}

func (str *MongoStore) GetRoute(ctx context.Context, id primitive.ObjectID) (*model.Route, error) {
	filter := bson.D{{"_id", id}}
	var route model.Route
	err := str.routes.FindOne(ctx, filter).Decode(&route)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoRoute)
		}
		return nil, err
	}
	return &route, err
}

func (str *MongoStore) GetFueler(ctx context.Context, id primitive.ObjectID) (*model.FuelingPerson, error) {
	filter := bson.D{{"_id", id}}
	var fueler model.FuelingPerson
	err := str.fuelers.FindOne(ctx, filter).Decode(&fueler)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoFueler)
		}
		return nil, err
	}
	return &fueler, err
}

func (str *MongoStore) GetFueling(ctx context.Context, id primitive.ObjectID) (*model.Fueling, error) {
	filter := bson.D{{"_id", id}}
	var fueling model.Fueling
	err := str.fuelings.FindOne(ctx, filter).Decode(&fueling)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoFueling)
		}
		return nil, err
	}
	return &fueling, err
}

func (str *MongoStore) GetMaintainer(ctx context.Context, id primitive.ObjectID) (*model.MaintenancePerson, error) {
	filter := bson.D{{"_id", id}}
	var maintainer model.MaintenancePerson
	err := str.mainPersons.FindOne(ctx, filter).Decode(&maintainer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoMaintainer)
		}
		return nil, err
	}
	return &maintainer, err
}

func (str *MongoStore) GetMaintenance(ctx context.Context, id primitive.ObjectID) (*model.Maintenance, error) {
	filter := bson.D{{"_id", id}}
	var maintenance model.Maintenance
	err := str.maintenances.FindOne(ctx, filter).Decode(&maintenance)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoMaintenance)
		}
		return nil, err
	}
	return &maintenance, err
}

func (str *MongoStore) GetVehicle(ctx context.Context, id primitive.ObjectID) (*model.Vehicle, error) {
	filter := bson.D{{"_id", id}}
	var vehicle model.Vehicle
	err := str.vehicles.FindOne(ctx, filter).Decode(&vehicle)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoVehicle)
		}
		return nil, err
	}
	return &vehicle, err
}

func (str *MongoStore) GetAuction(ctx context.Context, id primitive.ObjectID) (*model.Auction, error) {
	filter := bson.D{{"_id", id}}
	var auction model.Auction
	err := str.auctions.FindOne(ctx, filter).Decode(&auction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoAuction)
		}
		return nil, err
	}
	return &auction, err
}

func (str *MongoStore) GetReport(ctx context.Context, id primitive.ObjectID) (*model.Report, error) {
	filter := bson.D{{"_id", id}}
	var report model.Report
	err := str.reports.FindOne(ctx, filter).Decode(&report)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoReport)
		}
		return nil, err
	}
	return &report, err
}

func (str *MongoStore) GetRoutesByDriver(ctx context.Context, id primitive.ObjectID) ([]*model.Route, error) {
	filter := bson.D{{"driver_id", id}}
	res, err := str.routes.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var result []*model.Route
	err = res.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (str *MongoStore) GetActiveRoutesByDriver(ctx context.Context, id primitive.ObjectID) (*model.Route, error) {
	filter := bson.D{{"driver_id", id}, {"status", "active"}}
	var result model.Route
	err := str.routes.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(stat.NoReport)
		}
		return nil, err
	}
	return &result, err
}

func (str *MongoStore) CancelRoute(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{
		"status":    "active",
		"driver_id": id,
	}

	update := bson.M{
		"$set": bson.M{"status": "canceled"},
	}
	_, err := str.routes.UpdateMany(ctx, filter, update)
	return err
}

func (str *MongoStore) CompleteRoute(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{
		"status":    "active",
		"driver_id": id,
	}

	update := bson.M{
		"$set": bson.M{"status": "complete"},
	}
	_, err := str.routes.UpdateMany(ctx, filter, update)
	return err
}

func (str *MongoStore) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	filter2 := bson.M{"driver_id": id}
	filter3 := bson.M{"fueler_id": id}
	filter4 := bson.M{"maintenance_person_id": id}
	_, err := str.users.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	_, err = str.admins.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete admin: %v", err)
	}

	_, err = str.drivers.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete driver: %v", err)
	}

	_, err = str.fuelers.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete fueler: %v", err)
	}

	_, err = str.mainPersons.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete maintainer: %v", err)
	}

	_, err = str.routes.DeleteMany(ctx, filter2)
	if err != nil {
		return fmt.Errorf("failed to delete routes: %v", err)
	}

	_, err = str.maintenances.DeleteMany(ctx, filter4)
	if err != nil {
		return fmt.Errorf("failed to delete maintenances: %v", err)
	}

	_, err = str.fuelings.DeleteMany(ctx, filter3)
	if err != nil {
		return fmt.Errorf("failed to delete fuelings: %v", err)
	}

	_, err = str.vehicles.DeleteMany(ctx, filter2)
	if err != nil {
		return fmt.Errorf("failed to delete vehicles: %v", err)
	}

	_, err = str.reports.DeleteMany(ctx, filter2)
	if err != nil {
		return fmt.Errorf("failed to delete reports: %v", err)
	}
	return nil
}

func (str *MongoStore) DeleteDriver(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	filter2 := bson.M{"driver_id": id}
	_, err := str.users.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	_, err = str.admins.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete admin: %v", err)
	}

	_, err = str.drivers.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete driver: %v", err)
	}

	_, err = str.fuelers.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete fueler: %v", err)
	}

	_, err = str.mainPersons.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete maintainer: %v", err)
	}

	_, err = str.routes.DeleteMany(ctx, filter2)
	if err != nil {
		return fmt.Errorf("failed to delete routes: %v", err)
	}

	_, err = str.vehicles.DeleteMany(ctx, filter2)
	if err != nil {
		return fmt.Errorf("failed to delete vehicles: %v", err)
	}

	_, err = str.reports.DeleteMany(ctx, filter2)
	if err != nil {
		return fmt.Errorf("failed to delete reports: %v", err)
	}
	return nil
}

func (str *MongoStore) DeleteFueler(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	filter3 := bson.M{"fueler_id": id}
	_, err := str.users.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	_, err = str.admins.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete admin: %v", err)
	}

	_, err = str.drivers.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete driver: %v", err)
	}

	_, err = str.fuelers.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete fueler: %v", err)
	}

	_, err = str.mainPersons.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete maintainer: %v", err)
	}

	_, err = str.fuelings.DeleteMany(ctx, filter3)
	if err != nil {
		return fmt.Errorf("failed to delete fuelings: %v", err)
	}
	return nil
}

func (str *MongoStore) DeleteMaintainer(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	filter4 := bson.M{"maintenance_person_id": id}
	_, err := str.users.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	_, err = str.admins.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete admin: %v", err)
	}

	_, err = str.drivers.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete driver: %v", err)
	}

	_, err = str.fuelers.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete fueler: %v", err)
	}

	_, err = str.mainPersons.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete maintainer: %v", err)
	}

	_, err = str.maintenances.DeleteMany(ctx, filter4)
	if err != nil {
		return fmt.Errorf("failed to delete maintenances: %v", err)
	}
	return nil
}

func (str *MongoStore) DeleteVehicle(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	filter2 := bson.M{"vehicle_id": id}
	_, err := str.vehicles.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	_, err = str.fuelings.DeleteMany(ctx, filter2)
	if err != nil {
		return fmt.Errorf("failed to delete admin: %v", err)
	}

	_, err = str.maintenances.DeleteMany(ctx, filter2)
	if err != nil {
		return fmt.Errorf("failed to delete routes: %v", err)
	}

	_, err = str.auctions.DeleteMany(ctx, filter2)
	if err != nil {
		return fmt.Errorf("failed to delete maintenances: %v", err)
	}

	return nil
}

func (str *MongoStore) DeleteFueling(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := str.fuelings.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete fueling: %v", err)
	}

	return nil
}

func (str *MongoStore) DeleteMaintenance(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := str.maintenances.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete maintenance: %v", err)
	}

	return nil
}

func (str *MongoStore) DeleteRoutes(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	filter2 := bson.M{"route_id": id}
	_, err := str.routes.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete route: %v", err)
	}

	_, err = str.reports.DeleteMany(ctx, filter2)
	if err != nil {
		return fmt.Errorf("failed to delete admin: %v", err)
	}

	return nil
}

func (str *MongoStore) DeleteAuction(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := str.auctions.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete maintenance: %v", err)
	}

	return nil
}

func (str *MongoStore) DeleteReport(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := str.reports.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete maintenance: %v", err)
	}

	return nil
}
