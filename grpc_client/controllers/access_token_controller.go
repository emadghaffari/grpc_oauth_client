package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/emadghaffari/grpc_oauth_client/grpc_client/app"
	"github.com/emadghaffari/grpc_oauth_client/grpc_client/protos/accesstokenpb"
	"github.com/emadghaffari/res_errors/errors"
)



var(
	// AccessToken var
	AccessToken accessTokenInterface = &accessToken{}
	storeResponse *accesstokenpb.AccessTokenResponse
	getResponse *accesstokenpb.AccessTokenResponse
)


// accessTokenInterface interface
// the interface for accessTokens in dao
type accessTokenInterface interface{
	Get(string) (*accesstokenpb.AccessTokenResponse, error)
	Store(int32,int32) (*accesstokenpb.AccessTokenResponse, error)
	Delete(string) (*accesstokenpb.AccessTokenResponse, error)
	Update(string) (*accesstokenpb.AccessTokenResponse, error)
}

// accessToken struct implement all methods in interface
type accessToken struct{}

func (ac *accessToken)  Get(accessToken string) (*accesstokenpb.AccessTokenResponse, error){
	conn,err := app.StartApplication()
	if err != nil {
		return nil,err
	}
	defer conn.Close()
	req := &accesstokenpb.GetAccessTokenRequest{
		AccessToken: accessToken,
	}
	c :=  accesstokenpb.NewAccessTokenClient(conn)
	stream,err := c.Get(context.Background())
	wiatc := make(chan struct{})
	go func() {
		fmt.Printf("sending request ... %v \n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
		stream.CloseSend()
	}()

	go func() {
		res,err := stream.Recv()
		if err == io.EOF{
			close(wiatc)
			return
		}
		if err != nil {
			log.Fatalln(err.Error())
			close(wiatc)
			return
		}
		getResponse = res
	}()                             

	<-wiatc

	if storeResponse == nil {
		return nil, errors.HandlerInternalServerError("error in response section",nil)
	}

	return storeResponse, nil
}
func (ac *accessToken)  Store(userID int32, clientID int32) (*accesstokenpb.AccessTokenResponse, error){
	conn,err := app.StartApplication()
	if err != nil {
		return nil,err
	}
	defer conn.Close()
	req := &accesstokenpb.StoreAccessTokenRequest{
		ClientId: clientID,
		UserId: userID,
	}
	c :=  accesstokenpb.NewAccessTokenClient(conn)
	stream,err := c.Store(context.Background())
	wiatc := make(chan struct{})
	go func() {
		fmt.Printf("sending request ... %v \n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
		stream.CloseSend()
	}()

	go func() {
		res,err := stream.Recv()
		if err == io.EOF{
			close(wiatc)
			return
		}
		if err != nil {
			log.Fatalln(err.Error())
			close(wiatc)
			return
		}
		storeResponse = res
	}()

	<-wiatc

	if storeResponse == nil {
		return nil, errors.HandlerInternalServerError("error in response section",nil)
	}

	return storeResponse, nil
}

// Delete method for Delete accessTokens
func (ac *accessToken)  Delete(accessToken string) (*accesstokenpb.AccessTokenResponse, error){
	conn,err := app.StartApplication()
	if err != nil {
		return nil,err
	}
	defer conn.Close()
	c :=  accesstokenpb.NewAccessTokenClient(conn)
	req := &accesstokenpb.DeleteAccessTokenRequest{
		AccessToken: accessToken,
	}

	res,err := c.Delete(context.Background(),req)

	if err != nil {
		return nil , errors.HandlerInternalServerError("Error in Delete accessToken from Client service",err) 
	}
	return res, nil
}

// Update method for update accessTokens
func (ac *accessToken)  Update(accessToken string) (*accesstokenpb.AccessTokenResponse, error){
	conn,err := app.StartApplication()
	if err != nil {
		return nil,err
	}
	defer conn.Close()
	c :=  accesstokenpb.NewAccessTokenClient(conn)
	req := &accesstokenpb.UpdateAccessTokenRequest{
		AccessToken: accessToken,
	}

	res,err := c.Update(context.Background(),req)

	if err != nil {
		return nil , errors.HandlerInternalServerError("Error in Delete accessToken from Client service",err) 
	}
	return res, nil
}