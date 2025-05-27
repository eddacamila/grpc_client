// adapter/tren_adapter.go
package adapter

type TrenClientData struct {
	OpID      string
	Linea     string
	Capacidad int
	Estado    string
	Fecha     string
}

func (t TrenClientData) ToProto() *OperationalData {
	return &OperationalData{
		OperatorID:    t.OpID,
		RouteID:       t.Linea,
		Occupancy:     int32(t.Capacidad),
		VehicleStatus: t.Estado,
		Timestamp:     t.Fecha,
	}
}
