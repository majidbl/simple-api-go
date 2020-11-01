package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/majidzarephysics/university/api"
)

func TestIndexRoute(t *testing.T) {
    // The setupServer method, that we previously refactored
    // is injected into a test server
    router := gin.Default()
    router.GET("/ping", index)
    //api.ApplyRoutes(router)
    ts := httptest.NewServer(router)
    // Shut down the server and block until all requests have gone through
    defer ts.Close()

    // Make a request to our server with the {base url}/ping
    resp, err := http.Get(fmt.Sprintf("%s/ping", ts.URL))

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if resp.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
    }

    val, ok := resp.Header["Content-Type"]

    // Assert that the "content-type" header is actually set
    if !ok {
        t.Fatalf("Expected Content-Type header to be set")
    }

    // Assert that it was set as expected
    if val[0] != "application/json; charset=utf-8" {
        t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
    }
}
func TestApiRoute(t *testing.T) {
    // The setupServer method, that we previously refactored
    // is injected into a test server
    router := gin.Default()
    api.ApplyRoutes(router)
    ts := httptest.NewServer(router)
    // Shut down the server and block until all requests have gone through
    defer ts.Close()

    // Make a request to our server with the {base url}/ping
    resp, err := http.Get(fmt.Sprintf("%s/api/ping", ts.URL))

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if resp.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
    }

    val, ok := resp.Header["Content-Type"]

    // Assert that the "content-type" header is actually set
    if !ok {
        t.Fatalf("Expected Content-Type header to be set")
    }

    // Assert that it was set as expected
    if val[0] != "application/json; charset=utf-8" {
        t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
    }
}
func TestClassRoute(t *testing.T) {
    // The setupServer method, that we previously refactored
    // is injected into a test server
    router := gin.Default()
    api.ApplyRoutes(router)
    ts := httptest.NewServer(router)
    // Shut down the server and block until all requests have gone through
    defer ts.Close()

    // Make a request to our server with the {base url}/ping
    resp, err := http.Get(fmt.Sprintf("%s/api/v1.0/class/ping", ts.URL))

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if resp.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
    }

    val, ok := resp.Header["Content-Type"]

    // Assert that the "content-type" header is actually set
    if !ok {
        t.Fatalf("Expected Content-Type header to be set")
    }

    // Assert that it was set as expected
    if val[0] != "application/json; charset=utf-8" {
        t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
    }
}