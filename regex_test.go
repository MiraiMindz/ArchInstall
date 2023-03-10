package main

import (
	"fmt"
	"regexp"
	"testing"
)

//var regex string = "^(?=.{1,255}$)(?!.*\\.-|\\-.)[a-zA-Z0-9\\-]{1,63}(\\.[a-zA-Z0-9\\-]{1,63})*$"
//var regex string = "^(?=.{1,255}$)[0-9A-Za-z](?:(?:[0-9A-Za-z]|-){0,61}[0-9A-Za-z])?(?:\\.[0-9A-Za-z](?:(?:[0-9A-Za-z]|-){0,61}[0-9A-Za-z])?)*\\.?$"
//var regex string = "^[a-zA-Z0-9]([-a-zA-Z0-9]{0,61}[a-zA-Z0-9])?(\\.[a-zA-Z0-9]([-a-zA-Z0-9]{0,61}[a-zA-Z0-9])?)*$"
//var regex string = `^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`
var regex string = `^[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9](\.[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9])*$`

func TestValidHostnames(t *testing.T) {
	validHostnames := []string{
		"example.com",
		"subdomain.example.com",
		"subdomain.example.co.uk",
		"localhost",
		"www.example.com",
		"example_host.com",
		"example-host_1.example.com",
		"localhost",
		"myhost",
		"my-host",
		"my-host-1",
		"my_host",
		"myhost.local",
		"myhost.com",
		"www.example.com",
		"www.example.co.uk",
	}
	fmt.Println(regex)

	for _, hostname := range validHostnames {
		x, _ := regexp.MatchString(regex, hostname)
		if !x {
			t.Errorf("expected hostname %q to match regex, but it did not", hostname)
		}
	}
}

func TestInvalidHostnames(t *testing.T) {
	invalidHostnames := []string{
		"127.0.0.1",
		"_example.com",
		"example-.com",
		"example.com.",
		"example..com",
		"example_.com",
		"",
		"-myhost",
		"myhost-",
		"my_host_",
		"my..host",
		"my-host.",
		"myhost..com",
		"www.-example.com",
		"www.example-.com",
		"www.example.c0m",
		"123.123.123.123",
	}

	for _, hostname := range invalidHostnames {
		x, _ := regexp.MatchString(regex, hostname)
		if x {
			t.Errorf("expected hostname %q to not match regex, but it did", hostname)
		}
	}
}
