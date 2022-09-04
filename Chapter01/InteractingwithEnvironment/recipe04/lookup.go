package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	key := "DB_CONN"
	connStr, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("The env variable %s is not set.\n", key)
	}
	fmt.Println(connStr)
}
