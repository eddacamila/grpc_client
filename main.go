package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"grpc_client/adapter"
	pb "grpc_client/transport"

	"google.golang.org/grpc"
)

var estados = []string{"EnRuta", "Demorado", "FueraServicio", "Mantenimiento"}

func randomEstado() string {
	return estados[rand.Intn(len(estados))]
}

func randomOcupacion() int {
	return rand.Intn(100) // 0-99 pasajeros
}

func randomID(prefix string) string {
	return prefix + "-" + string('A'+rune(rand.Intn(26)))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	conn, err := grpc.Dial("ec2-3-128-197-20.us-east-2.compute.amazonaws.com:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewTransportServiceClient(conn)

	for i := 0; i < 35; i++ {
		var adapted *adapter.OperationalData

		if rand.Intn(2) == 0 {
			// Cliente tipo Bus
			dataBus := adapter.BusClientData{
				CodOperador: randomID("BUS"),
				Ruta:        randomID("R"),
				Ocupacion:   randomOcupacion(),
				EstadoBus:   randomEstado(),
				Tiempo:      time.Now().Format(time.RFC3339),
			}
			adapted = dataBus.ToProto()
		} else {
			// Cliente tipo Tren
			dataTren := adapter.TrenClientData{
				OpID:      randomID("TRN"),
				Linea:     randomID("L"),
				Capacidad: randomOcupacion(),
				Estado:    randomEstado(),
				Fecha:     time.Now().Format(time.RFC3339),
			}
			adapted = dataTren.ToProto()
		}

		res, err := client.SendOperationalData(context.Background(), &pb.OperationalData{
			OperatorId:    adapted.OperatorID,
			RouteId:       adapted.RouteID,
			Occupancy:     adapted.Occupancy,
			VehicleStatus: adapted.VehicleStatus,
			Timestamp:     adapted.Timestamp,
		})
		if err != nil {
			log.Printf("❌ Error al enviar datos [%d]: %v", i, err)
		} else {
			log.Printf("✅ [%d] Enviado a gRPC → %s", i, res.Message)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
