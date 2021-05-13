# Multiplica

A GRPC & API Rest service to provide a single operation: multiplication of 2 numbers.

## Notes

I have opted for scoping the range of values in float32. Implementation could be more easy if I use float64 because it's the standard of grpc but, as a way of saving space.
Another option would be using the go package big (Big Numbers) which wraps very well math operations without losing precision on long values. A drawback of this is that is computationally more expensive.
Because of these reasons, for this escenario, I opted for float32 value type as a good choice for this implementation considering the use of values won't be extra large numbers (exceeding float32).   

CI/CD: Github Actions (✔️)
Containerized (✔️)

## What this service does?

Multiply 2 numbers. The number values ranges all IEEE-754 32-bit floating-point numbers (about 7 decimal digits, and a max value of 3.40282346638528859811704183484516925440e38).

If the 2 numbers are big, then multiplication could lead to an "infinite" number, which will lead to return an error.

Incorrect parameters (such as missing one parameter) will lead to return an error also.

## Grpc

### How to Use

- Go way

In one terminal, run this command to initiate server:

        go run cmd/grpc-server/main.go

In other terminal, run this command to initiate make a call as client and make the operation 10 times 3:

        go run cmd/grpc-client/main.go -x 10 -y 3

- Binary way (First execute build.sh to generate the builds):

In one terminal, run this command to initiate server:

        ./grpc-server

In other terminal, run this command to initiate make a call as client and make the operation 10 times 3:

        ./grpc-client -x 10 -y 3

### Docs

Further details of GRPC API are located on docs folder


## Rest

REST API service. Is not limited to a specific METHOD (So is possible to use on any request Method: GET, POST, PUT, etc etc)

### How to Use

- Go way

        go run cmd/rest/main.go

- Binary way (First execute build.sh to generate the builds):

In one terminal, run this command to initiate server (default port: 11000):

        ./rest-server

In other terminal, run this command to make a call to the rest service:

        curl -X GET "http://localhost:11000/v1/multiply?x=23&y=4343"


## TODO

Some steps were excepted but for a production release these steps should be considered:
- Security (Serving services by a security layer, REST: serve on https, GRPC: add TLS certificates)
- Telemetry. Promotheus is a go-to telemetry tools well known and useful to instrument services 
- An advanced CI/CD Attaching to a container storage to deploy and store generated builds 
