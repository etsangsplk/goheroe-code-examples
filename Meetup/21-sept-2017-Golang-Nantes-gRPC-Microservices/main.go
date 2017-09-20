func main() {
	
		lis, err := net.Listen("tcp", grpcPort)
		if err != nil {
			fmt.Printf("failed to listen: %v\n", err)
		}
		
		grpcServer := grpc.NewServer()

		srv, err := app.NewGoHeroeServer()
		if err != nil {
			fmt.Printf("failed to create GoHeroeServer : %s\n", err)
		}
		
		superpower.RegisterGoHeroeServer(grpcServer, srv)
	
		reflection.Register(grpcServer)
	
		fmt.Printf("starting server on port : %s\n", grpcPort)
		grpcServer.Serve(lis)
	}