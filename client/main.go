package main

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"

	"github.com/Selahattinn/grpc-weather/api"
)

func main() {
	addr := "localhost:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := api.NewWeatherServiceClient(conn)

	ctx := context.Background()

	resp, err := client.ListCities(ctx, &api.ListCitiesRequest{})
	if err != nil {
		panic(err)
	}

	for _, city := range resp.Items {
		fmt.Printf("City: %s, Code: %d\n", city.GetName(), city.GetCode())
	}

	stream, err := client.QueryWeather(ctx, &api.WeatherRequest{
		CityCode: 35,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Weather stream started")
	fmt.Println("Weather in Izmir:")
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("Temperature: ", resp.GetTemperature())
	}

	fmt.Println("Server closed the stream")
}
