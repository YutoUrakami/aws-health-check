package main

import (
	"context"
	"fmt"

	"./healthevent"
	"./slack"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambdaEvent(ctx context.Context, event healthevent.Event) (string, error) {
	fmt.Println(event.Detail.Service)
	fmt.Println("  [TypeCode]: " + event.Detail.TypeCode)
	fmt.Println("  [TypeCategory]: " + event.Detail.TypeCategory)
	fmt.Println("  [StartTime]: " + event.Detail.StartTime)
	fmt.Println("  [EndTime]: " + event.Detail.EndTime)
	fmt.Println("  [Description]: " + event.Detail.Description[0].Latest)

	err := slack.Send(&event.Detail)
	if err != nil {
		fmt.Print("!!!!![ERROR]!!!!! : " + err.Error())
	}

	return "", err
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
