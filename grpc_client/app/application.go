package app

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/emadghaffari/res_errors/errors"
)

const(
	certFile = "ssl/server.crt"
	keyFile = "ssl/server.pem"
)

// StartApplication func
// starter for application
func StartApplication() (*grpc.ClientConn ,error){
	// if go code crashed...
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	/* SSL */
	// certs,err := credentials.NewClientTLSFromFile(certFile,keyFile)
	// if err != nil {
	// 	fmt.Println(errors.HandlerInternalServerError(fmt.Sprintf("Error in credential Client TLS from File: %v", err),err))
	// 	return
	// }
	// conn,err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	
	
	conn,err := grpc.Dial(":50051", grpc.WithInsecure())
	
	if err != nil {
		return nil, errors.HandlerInternalServerError(fmt.Sprintf("Error in Dial to grpc server: %v", err),err)
	}
	
	return conn,nil

}