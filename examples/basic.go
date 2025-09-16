package main

import (
	"fmt"
	"log"

	"github.com/nurettintopal/semaver"
)

func main() {
	result, err := semaver.CompareVersions("1.0.0-alpha.1", "<", "1.0.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha.1", ">", "1.0.0-alpha")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha.1", "=", "1.0.0-alpha.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-beta.2", "<", "1.0.0-beta.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0+build123", ">", "1.0.0+build456")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0", "<", "1.2.2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.2.1", ">", "1.2.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("2.0.3", "=", "2.0.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("2.0.2", ">=", "2.0.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha.1", "<", "1.0.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha.1", ">", "1.0.0-alpha")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha.1", "=", "1.0.0-alpha.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-beta.2", "<", "1.0.0-beta.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.1.0-beta.2", "<", "1.0.0-beta.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha", ">=", "1.0.0-alpha.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0+build123", ">", "1.0.0+build456")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha+build789", "<", "1.0.0-alpha+build000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.2.0-alpha", "<", "1.2.0-beta")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.2.0-alpha.1", ">", "1.2.0-alpha")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0+build0001", "=", "1.0.0+build0001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0", "<", "1.2.2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.2.1", ">", "1.2.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("2.0.3", "=", "2.0.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("2.0.2", ">=", "2.0.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.4.3", ">=", "1.0.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0", "<=", "1.0.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0", ">", "1.0.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)
}
