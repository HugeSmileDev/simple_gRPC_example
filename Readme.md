# gRPC Client-Server Application in Golang

This project is a simple client-server application that communicates using gRPC. It is written in Golang and allows you to send your name to the server and receive a greeting in response.

## Prerequisites

Before running this project, make sure you have the following installed:

- Go programming language (https://golang.org/doc/install)

## Getting Started

1. Clone this Git repository to your local machine:

```bash
git clone <repository-url>
```

2. Navigate to the project directory:

```bash
cd <project-directory>
```

3. Compile the proto file to generate gRPC service code:

```bash
protoc --go_out=. --go-grpc_out=. tutorial/tutorial.proto
```

## Running the Application

Follow these steps to run the client-server application:

1. Start the server by running the `main.go` file:

```bash
go run main.go
```

2. In a separate terminal window, run the client application using the `client.go` file:

```bash
go run client.go
```

3. Once the client is running, it will prompt you to enter your name. Type your name and press `Enter`.

4. The server will receive your name, generate a greeting message, and send it back to the client.

5. The client will then display the received greeting message.

## Project Structure

The project consists of the following files:

- `main.go`: This file contains the code for the server application. It sets up the gRPC server and handles client requests.

- `client.go`: This file contains the code for the client application. It connects to the server, sends a name to the server, and receives the greeting response.

- `tutorial/tutorial.proto`: This file is a Protocol Buffer file (proto) defining the gRPC service and messages used for communication between the client and server.

## gRPC Service Definition

The `tutorial/tutorial.proto` file defines the following gRPC service and messages:

```proto
syntax = "proto3";

package tutorial;

service Tutorial {
    rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
    string name = 1;
}

message HelloReply {
    string message = 1;
}
```

- `Tutorial`: This is the gRPC service that defines a single RPC method `SayHello`. When the client calls this method with a `HelloRequest`, the server responds with a `HelloReply`.

- `HelloRequest`: This message contains a single field `name`, which represents the client's name.

- `HelloReply`: This message contains a single field `message`, which represents the greeting message sent by the server.

## Conclusion

Congratulations! You have successfully set up and run the gRPC client-server application in Golang. Feel free to explore the code and experiment with gRPC communication between the client and server. If you encounter any issues or have suggestions for improvement, please don't hesitate to open an issue or contribute to the project. Happy coding!