# gRPC

## Reflection service

We register our server with the reflection service (`reflection.Register(s)`) so that the service provides information about the publicly accessible services on the server.

This is not to be used if you are requiring the client accessing the server to have the proto definition.

The idea is that it makes the service dynamically callable. All the possible options a client can use can be seen to them. They would not need to reference the originally proto file for the server.
