package init

import (
	"api/basic/global"
	__ "api/basic/proto"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func init() {
	InitGrpc()
}

func InitGrpc() {

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	global.UserClient = __.NewUserClient(conn)

}
