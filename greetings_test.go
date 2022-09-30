package greetings

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	message := fmt.Sprintf("Hello world, my name is %v !", t.Name())
	fmt.Println(message)
}
