package main

import (
	"context"
	observabilityGo "github.com/rzeAkbari/observabilityGo/server/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewClient() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := observabilityGo.NewRouteGuideClient(conn)

	feature, err := client.GetFeature(context.Background(),
		&observabilityGo.Point{Latitude: 409146138, Longitude: -746188906})
	log.Println(feature)
}

func main() {
	NewClient()
}
