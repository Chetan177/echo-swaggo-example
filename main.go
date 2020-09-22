package main

import (
	_ "echo-swaggo-example/docs"
	"echo-swaggo-example/pkg/restserver"
)



func main()  {
	rest := restserver.Rest{Port:"8443"}
	rest.StartServer()

}
