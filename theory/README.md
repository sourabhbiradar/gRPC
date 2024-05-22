# gRPC framework

- Advanced version of RPC developed by Google.
- Uses `HTTP/2` 
- `Protocol Buffers` (Protobuf) as the interface description language
- For Load Balancing uses `Nginx`

- Supports 15+ languages.

`RPC` : Remote Procedure Call
`gRPC` : `g` represents kinda version , its different with every new release.

`Procedure` : Function or method that is called by `Client Program` but is executed on `Remote Server`.

Abstrating network communication making it appear as if the procedure is executed locally.

HOW gRPC WORKS ?
1) `Client Stub` : client makes local procedure call that invokes RPC. 
This call is handled by `Client Stub` which packs parameters into message & sends to Server.
It `marshals` params & sends.

2) `Communicatoin` : request-response are received-sent over `network`

3) `Server Stub`  : receives request, unpacks parameters & invokes appropriate procedure (function) on server.
It `unmarshals` params --> invoke procedure --> `marshal` response.

4) `Execution` : server executes request & sends back response.

5) `Response Handling` : `Server Stub` packs response, sends it back to client.

6) `Client Handling` : `Client Stub` receives the response, unpacks it, returns it to `client Application`.

NOTE : `Stubs` = `Messages` = `protobuf file` (.proto file / proto buffer file)


![alt text][def]

- json , xml vs buffer :
json   : {"key" : "value"} --> light wieght

xml    : <tag> </tag>    --> Heavy wieght

buffer : [bytes] >> array of bytes --> lighter

- Types of gRPC/RPC :
1) `Unary RPC` is the simplest form of gRPC where the client sends a single request to the server and receives a single response.
2) `Server streaming RPC`: the client sends a single request to the server and receives a stream of responses.
3) `Client streaming RPC`: the client sends a stream of requests to the server and receives a single response.
4) `Bidirectional streaming RPC` : both the client and the server send a sequence of messages using a read-write stream. The two streams operate independently, so clients and servers can read and write in whatever order they like.









[def]: image.png
