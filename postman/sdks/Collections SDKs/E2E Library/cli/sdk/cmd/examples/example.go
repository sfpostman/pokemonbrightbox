package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"context"
	"example.com/e2e-library/sdk"
	"example.com/e2e-library/sdk/books"
)

func main() {
	loadEnv()

	config := e2elibrarysdk.NewConfig()
	client := e2elibrarysdk.NewE2eLibrarySDK(config)

	params := books.FetchAListOfBooksRequestParams{
		XMockResponseCode: e2elibrarysdk.Nullable[string]("500"),
		Accept:            e2elibrarysdk.Nullable[string]("application/json"),
	}

	response, err := client.Books.FetchAListOfBooks(context.Background(), params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", response)
}

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
