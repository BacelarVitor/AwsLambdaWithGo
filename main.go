package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"what is your name?"`
	Age  int    `json:"How old are you?"`
}

type MyReponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyReponse, error) {
	return MyReponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
