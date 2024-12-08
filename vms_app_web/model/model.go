package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Email       string             `bson:"email" json:"email"`
	FirstName   string             `bson:"first_name" json:"first_name"`
	LastName    string             `bson:"last_name" json:"last_name"`
	MiddleName  string             `bson:"middle_name" json:"middle_name"`
	Password    string             `bson:"password" json:"password"`
	RoleID      primitive.ObjectID `bson:"role_id" json:"role_id"`
	GovermentID int                `bson:"goverment_id" json:"goverment_id"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number"`
}

func NewUser(Email string, FirstName string, LastName string, MiddleName string, Password string, RoleID primitive.ObjectID, GovermentID int, PhoneNumber string) *User {
	return &User{
		ID:          primitive.NewObjectID(),
		Email:       Email,
		FirstName:   FirstName,
		LastName:    LastName,
		MiddleName:  MiddleName,
		Password:    Password,
		RoleID:      RoleID,
		GovermentID: GovermentID,
		PhoneNumber: PhoneNumber,
	}
}

type Role struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Permissions int                `bson:"permissions" json:"permissions"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

func NewRole(Name string, Description string, Permissions int) *Role {
	return &Role{
		ID:          primitive.NewObjectID(),
		Name:        Name,
		Description: Description,
		Permissions: Permissions,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

type Admin struct {
	ID primitive.ObjectID `bson:"_id"`
}

func NewAdmin(userID primitive.ObjectID) *Admin {
	return &Admin{
		ID: userID,
	}
}

type Driver struct {
	ID     primitive.ObjectID `bson:"_id"`
	Adress string             `bson:"adress"`
}

func NewDriver(userID primitive.ObjectID, Adress string) *Driver {
	return &Driver{
		ID:     userID,
		Adress: Adress,
	}
}

type Route struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	DriverID      primitive.ObjectID `bson:"driver_id" json:"driver_id"`
	StartLocation string             `bson:"start_location" json:"start_location"`
	EndLocation   string             `bson:"end_location" json:"end_location"`
	Description   string             `bson:"description" json:"description"`
	TimeForRoute  int                `bson:"time_for_route" json:"time_for_route"`
	Distance      float32            `bson:"distance" json:"distance"`
	Status        string             `bson:"status" json:"status"`
}

func NewRoute(DriverID primitive.ObjectID, StartLocation string, EndLocation string,
	Description string, TimeForRoute int, Distance float32, Status string) *Route {
	return &Route{
		ID:            primitive.NewObjectID(),
		DriverID:      DriverID,
		StartLocation: StartLocation,
		EndLocation:   EndLocation,
		Description:   Description,
		TimeForRoute:  TimeForRoute,
		Distance:      Distance,
		Status:        Status,
	}
}

type FuelingPerson struct {
	ID primitive.ObjectID `bson:"_id"`
}

func NewFuelingPerson(userID primitive.ObjectID) *FuelingPerson {
	return &FuelingPerson{
		ID: userID,
	}
}

type Fueling struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	FuelerID     primitive.ObjectID `bson:"fueler_id" json:"fueler_id"`
	VehicleID    string             `bson:"vehicle_id" json:"vehicle_id"`
	FuelAmount   float32            `bson:"fuel_amount" json:"fuel_amount"`
	TotalCost    float32            `bson:"total_cost" json:"total_cost"`
	Date         time.Time          `bson:"date" json:"date"`
	GasStation   string             `bson:"gas_station" json:"gas_station"`
	FuelingProof string             `bson:"fueling_proof" json:"fueling_proof"`
}

func NewFueling(FuelerID primitive.ObjectID, VehicleID string, FuelAmount float32,
	TotalCost float32, Date time.Time, GasStation string, FuelingProof string) *Fueling {
	return &Fueling{
		ID:           primitive.NewObjectID(),
		FuelerID:     FuelerID,
		VehicleID:    VehicleID,
		FuelAmount:   FuelAmount,
		TotalCost:    TotalCost,
		Date:         Date,
		GasStation:   GasStation,
		FuelingProof: FuelingProof,
	}
}

type MaintenancePerson struct {
	ID primitive.ObjectID `bson:"_id"`
}

func NewMaintenancePerson(userID primitive.ObjectID) *MaintenancePerson {
	return &MaintenancePerson{
		ID: userID,
	}
}

type Maintenance struct {
	ID                  primitive.ObjectID `bson:"_id" json:"_id"`
	MaintenancePersonID primitive.ObjectID `bson:"maintenance_person_id" json:"maintenance_person_id"`
	VehicleID           string             `bson:"vehicle_id" json:"vehicle_id"`
	Date                time.Time          `bson:"date" json:"date"`
	Parts               []string           `bson:"parts" json:"parts"`
	ServiceType         string             `bson:"service_type" json:"service_type"`
	Proof               string             `bson:"maintenance_proof" json:"maintenance_proof"`
	Cost                float32            `bson:"cost" json:"cost"`
}

func NewMaintenance(MaintenancePersonID primitive.ObjectID, VehicleID string, Parts []string,
	ServiceType string, Date time.Time, Proof string, Cost float32) *Maintenance {
	return &Maintenance{
		ID:                  primitive.NewObjectID(),
		MaintenancePersonID: MaintenancePersonID,
		VehicleID:           VehicleID,
		Parts:               Parts,
		ServiceType:         ServiceType,
		Date:                Date,
		Cost:                Cost,
		Proof:               Proof,
	}
}

type Vehicle struct {
	VIN            string             `bson:"_id" json:"vin"`
	ProductionYear int                `bson:"production_year" json:"production_year"`
	Mileage        int                `bson:"mileage" json:"mileage"`
	LicensePlate   string             `bson:"license_plate" json:"license_plate"`
	DriverID       primitive.ObjectID `bson:"driver_id" json:"driver_id"`
	Model          string             `bson:"model" json:"model"`
	CarMake        string             `bson:"car_make" json:"car_make"`
	ActivityStatus string             `bson:"activity_status" json:"activity_status"`
	ExactLocation  string             `bson:"exact_location" json:"exact_location"`
}

func NewVehicle(VIN string, ProductionYear int, Mileage int,
	LicensePlate string, DriverID primitive.ObjectID, Model string,
	CarMake string, ActivityStatus string, ExactLocation string) *Vehicle {
	return &Vehicle{
		VIN:            VIN,
		ProductionYear: ProductionYear,
		Mileage:        Mileage,
		LicensePlate:   LicensePlate,
		DriverID:       DriverID,
		Model:          Model,
		CarMake:        CarMake,
		ActivityStatus: ActivityStatus,
		ExactLocation:  ExactLocation,
	}
}

type Auction struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	VehicleID     string             `bson:"vehicle_id" json:"vehicle_id"`
	Name          string             `bson:"name" json:"name"`
	Description   string             `bson:"description" json:"description"`
	StartingPrice int                `bson:"starting_price" json:"starting_price"`
	CurrentPrice  int                `bson:"current_price" json:"current_price"`
	Status        string             `bson:"status" json:"status"`
}

func NewAuction(VehicleID string, Name string, Description string,
	StartingPrice int, CurrentPrice int, Status string) *Auction {
	return &Auction{
		ID:            primitive.NewObjectID(),
		VehicleID:     VehicleID,
		Name:          Name,
		Description:   Description,
		StartingPrice: StartingPrice,
		CurrentPrice:  CurrentPrice,
		Status:        Status,
	}
}

type Report struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	TotalDistance int                `bson:"total_distance" json:"total_distance"`
	MoneySpent    int                `bson:"money_spent" json:"money_spent"`
	FuelUsage     float32            `bson:"fuel_usage" json:"fuel_usage"`
	DriverID      primitive.ObjectID `bson:"driver_id" json:"driver_id"`
	VehicleID     string             `bson:"vehicle_id" json:"vehicle_id"`
	RouteID       primitive.ObjectID `bson:"route_id" json:"route_id"`
}

func NewReport(TotalDistance int, MoneySpent int, FuelUsage float32,
	DriverID primitive.ObjectID, VehicleID string, RouteID primitive.ObjectID) *Report {
	return &Report{
		ID:            primitive.NewObjectID(),
		TotalDistance: TotalDistance,
		MoneySpent:    MoneySpent,
		FuelUsage:     FuelUsage,
		DriverID:      DriverID,
		VehicleID:     VehicleID,
		RouteID:       RouteID,
	}
}
