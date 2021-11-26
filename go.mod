module github.com/kumoroku/sumologic-go-sdk

go 1.17

require github.com/kumoroku/go-sumologic v0.0.0

require (
	github.com/golang/protobuf v1.4.2 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/oauth2 v0.0.0-20210218202405-ba52d332ba99 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/kumoroku/go-sumologic => ../out/go-sumologic
