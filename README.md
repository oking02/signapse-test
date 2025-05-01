# signapse-test

## Information

> **Note**  
> This project layout is **definitely** overkill for just this functionality.
> The idea was to showcase a hexagonal architecture like structure that:
> - Makes extending the service much easier
> - Reduces friction for another engineer to extend
> 
> I haven't used strict hexagonal naming as I'm more used to using handlers/datasources vs adapters
> 
> Lots of boilerplate and code was taken from a new service template repo I've created.

I've tried to keep dependencies to the minimum but might use a frameworks for http in a larger services, as they provide lots of nice to haves and routing options

<img src="./arch1.png" width="800" alt="Medium">


## To Use

Run

`go run github.com/oking02/signapse-test/cmd/app`

This will start a server on localhost:3000. 
To change the port set the `HTTP_PORT` enviroment variable
