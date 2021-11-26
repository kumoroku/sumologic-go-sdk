package main

import (
	"context"
	"fmt"
	sumo "github.com/kumoroku/go-sumologic"
	"os"
)

func main() {

	configuration := sumo.NewConfiguration()
	configuration.Host = os.Getenv("SUMOLOGIC_HOST")
	configuration.Debug = true
	client := sumo.NewAPIClient(configuration)

	accessId := os.Getenv("SUMOLOGIC_ID")
	accessKey := os.Getenv("SUMOLOGIC_KEY")
	ctx := context.WithValue(context.Background(), sumo.ContextBasicAuth,
		sumo.BasicAuth{UserName: accessId, Password: accessKey})

	request := client.UserManagementApi.ListUsers(ctx)
	users, _, err := client.UserManagementApi.ListUsersExecute(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%v\n\n", users)
}

//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
