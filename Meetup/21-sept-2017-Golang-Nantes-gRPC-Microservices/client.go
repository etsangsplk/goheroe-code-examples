func main() {
	conn, err := grpc.Dial("localhost:44444", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	ctx := context.Background()
	client := superpower.NewGoHeroeClient(conn)
	filter := superpower.Filter{Category: superpower.SuperPowerCategory_CosmicBasedPowers}

	result, err := client.List(ctx, &filter)
	if err != nil {
		log.Fatalf("Error on List : %v", err)
	}
	log.Printf("List: %v", result)

	power := superpower.SuperPower{Name: "mickey power", Cat: 0, CoolPow: true}

	result, err = client.Add(ctx, &power)
	if err != nil {
		log.Fatalf("Error on List : %v", err)
	}
	log.Printf("List: %v", result)
}