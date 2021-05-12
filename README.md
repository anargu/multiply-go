# Multiplica

A GRPC & API Rest service to provide a single operation: multiplication of 2 numbers

## What this service does?

Multiply 2 numbers. The number values ranges all IEEE-754 32-bit floating-point numbers (about 7 decimal digits, and a max value of 3.40282346638528859811704183484516925440e38).


If the 2 numbers are big, then multiplication could lead to an "infinite" number, which will lead to return an error. 

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


## Rest

### How to Use

- Go way

        go run cmd/rest/main.go

- Binary way (First execute build.sh to generate the builds):

In one terminal, run this command to initiate server (default port: 11000):

        ./rest-server

In other terminal, run this command to make a call to the rest service:

        curl -X GET "http://localhost:11000/v1/multiply?x=23&y=4343"
