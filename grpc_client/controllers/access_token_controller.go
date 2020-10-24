package controllers

import (
	"context"
	"fmt"

	"github.com/emadghaffari/grpc_oauth_client/grpc_client/app"
	"github.com/emadghaffari/grpc_oauth_client/grpc_client/protos/accesstokenpb"
	"github.com/emadghaffari/res_errors/errors"
)



var(
	// ClientAccessToken var
	ClientAccessToken accessTokenInterface = &accessToken{}
	storeResponse *accesstokenpb.AccessTokenResponse
	getResponse *accesstokenpb.AccessTokenResponse
)


// accessTokenInterface interface
// the interface for accessTokens in dao
type accessTokenInterface interface{
	Get(string) (*accesstokenpb.AccessTokenResponse, errors.ResError)
	Store(int32,int32) (*accesstokenpb.AccessTokenResponse, errors.ResError)
	Delete(string) (*accesstokenpb.AccessTokenResponse, errors.ResError)
	Update(string) (*accesstokenpb.AccessTokenResponse, errors.ResError)
}

// accessToken struct implement all methods in interface
type accessToken struct{}

func (ac *accessToken)  Get(accessToken string) (*accesstokenpb.AccessTokenResponse, errors.ResError){
	conn,err := app.StartApplication()
	if err != nil {
		return nil,err
	}
	defer conn.Close()
	req := &accesstokenpb.GetAccessTokenRequest{
		AccessToken: accessToken,
	}
	c :=  accesstokenpb.NewAccessTokenClient(conn)
	res,reserr := c.Get(context.Background(),req)
	if reserr != nil {
		return nil, errors.HandlerInternalServerError(fmt.Sprintf("Error in Get data with accessToken from Client service"),reserr)
	}
	return res, nil
}
func (ac *accessToken)  Store(userID int32, clientID int32) (*accesstokenpb.AccessTokenResponse, errors.ResError){
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
	res,reserr := c.Store(context.Background(),req)
	if reserr != nil {
		return nil, errors.HandlerInternalServerError(fmt.Sprintf("Error in Store accessToken from Client service"),reserr)
	}
	return res, nil
}

// Delete method for Delete accessTokens
func (ac *accessToken)  Delete(accessToken string) (*accesstokenpb.AccessTokenResponse, errors.ResError){
	conn,err := app.StartApplication()
	if err != nil {
		return nil,err
	}
	defer conn.Close()
	c :=  accesstokenpb.NewAccessTokenClient(conn)
	req := &accesstokenpb.DeleteAccessTokenRequest{
		AccessToken: accessToken,
	}

	res,reserr := c.Delete(context.Background(),req)
	if reserr != nil {
		return nil, errors.HandlerInternalServerError(fmt.Sprintf("Error in Delete accessToken from Client service"),reserr)
	}
	return res, nil
}

// Update method for update accessTokens
func (ac *accessToken)  Update(accessToken string) (*accesstokenpb.AccessTokenResponse, errors.ResError){
	conn,err := app.StartApplication()
	if err != nil {
		return nil,err
	}
	defer conn.Close()
	c :=  accesstokenpb.NewAccessTokenClient(conn)
	req := &accesstokenpb.UpdateAccessTokenRequest{
		AccessToken: accessToken,
	}

	res,reserr := c.Update(context.Background(),req)

	if reserr != nil {
		return nil , errors.HandlerInternalServerError("Error in Delete accessToken from Client service",reserr) 
	}
	return res, nil
}