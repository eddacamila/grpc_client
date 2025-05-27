// adapter/adapter.go
package adapter

type OperationalDataInput interface {
	ToProto() *OperationalData
}

type OperationalData struct {
	OperatorID    string
	RouteID       string
	Occupancy     int32
	VehicleStatus string
	Timestamp     string
}
