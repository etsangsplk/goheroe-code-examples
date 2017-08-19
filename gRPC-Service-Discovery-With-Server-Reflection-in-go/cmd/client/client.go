package main

import (
	"log"

	"github.com/chemidy/goheroe-code-examples/gRPC-Service-Discovery-With-Server-Reflection-in-go/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:44444", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()

	// Creates a new NewGoHeroeClient
	client := superpower.NewGoHeroeClient(conn)

	//filter := superpower.Filter{Category: superpower.SuperPowerCategory_All}

	//filter := superpower.Filter{Category: superpower.SuperPowerCategory_SciencePowers}

	filter := superpower.Filter{Category: superpower.SuperPowerCategory_CosmicBasedPowers}

	// calling the  API
	result, err := client.List(ctx, &filter)
	if err != nil {
		log.Fatalf("Error on List : %v", err)
	}

	log.Printf("List: %v", result)

	power := superpower.SuperPower{Name: "mickey power", Cat: 0, CoolPow: true}

	// calling the  API
	result, err = client.Add(ctx, &power)
	if err != nil {
		log.Fatalf("Error on List : %v", err)
	}

	log.Printf("List: %v", result)

}
