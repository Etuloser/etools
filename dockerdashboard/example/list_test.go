package example

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	got, err := List()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(got)
}

func TestGetPublicIp(t *testing.T) {
	got, err := getPublicIp()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(got)
}
