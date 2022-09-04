package main

import (
	"log"
	"os"
)

func main() {
	key := "DB_CONN"
	// Set the environmental variable.
	err := os.Setenv(key, "postgres://as:as@example.com/pg?sslmode=verify-full")
	if err != nil {
		panic(err)
	}

	val := GetEnvDefault(key, "postgres://as:as@localhost/pg?sslmode=verify-full")

	log.Println("The value is: " + val)

	err = os.Unsetenv(key)
	if err != nil {
		panic(err)
	}

	val = GetEnvDefault(key, "postgres://as:as@127.0.0.1/pg?sslmode=verify-full")

	log.Println("The default value is: " + val)

}

func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
