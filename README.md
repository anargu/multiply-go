# Multiplica

A GRPC & API Rest service to provide a single operation: multiplication of 2 numbers

## What this service does?

Multiply 2 numbers. The number values ranges all IEEE-754 32-bit floating-point numbers (about 7 decimal digits, and a max value of 3.40282346638528859811704183484516925440e38).


If the 2 numbers are big, then multiplication could lead to an "infinite" number, which will lead to return an error. 

## Additional Details

### Re-generate proto files

Requeriments
- protoc already installed

