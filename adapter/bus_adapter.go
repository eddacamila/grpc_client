// adapter/client_bus.go
package adapter

type BusClientData struct {
	CodOperador string
	Ruta        string
	Ocupacion   int
	EstadoBus   string
	Tiempo      string
}

func (b BusClientData) ToProto() *OperationalData {
	return &OperationalData{
		OperatorID:    b.CodOperador,
		RouteID:       b.Ruta,
		Occupancy:     int32(b.Ocupacion),
		VehicleStatus: b.EstadoBus,
		Timestamp:     b.Tiempo,
	}
}
