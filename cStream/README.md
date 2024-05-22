# client streaming RPC

`IMPORTANT NOTE` : 
1) service implementation is diffenrent (ServerReply) ,go through it.
ServerReply() takes stream of requests & sends only one response

2) clientConnectionServer() (client.go) now sends stream of requests

go run server.go
go run client.go (on another terminal window) 

if no error , copy paste below line in browser
http://localhost:8008/sent


