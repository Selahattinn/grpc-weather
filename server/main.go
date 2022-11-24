package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/Selahattinn/grpc-weather/api"
)

func main() {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	api.RegisterWeatherServiceServer(srv, &myWeatherService{})

	fmt.Println("Starting server  on port 8080")
	panic(srv.Serve(list))
}

type myWeatherService struct {
	api.UnimplementedWeatherServiceServer
}

func (m *myWeatherService) ListCities(ctx context.Context, req *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {
	return &api.ListCitiesResponse{
		Items: []*api.City{
			{
				Name: "Istanbul",
				Code: 35,
			},
			{
				Name: "Ankara",
				Code: 6,
			},
			{
				Name: "Izmir",
				Code: 35,
			},
		},
	}, nil
}

func (m *myWeatherService) QueryWeather(req *api.WeatherRequest, resp api.WeatherService_QueryWeatherServer) error {
	for {
		if req.GetCityCode() != 35 {
			fmt.Println("City not found")
			break
		}
		err := resp.Send(&api.WeatherResponse{
			Temperature: rand.Float32() + 10,
		})
		if err != nil {
			fmt.Println("Error while sending response: ", err)
			break
		}

		time.Sleep(time.Second)
	}
	return nil
}
