package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const QueueURL = "http://sqs.us-west-2.localhost.localstack.cloud:4566/000000000000/my-first-queue"

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1"),
	}))

	svc := sqs.New(sess)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		select {
		case <-signalCh:
			fmt.Println("Saindo")
			return
		default:
			receiveParams := &sqs.ReceiveMessageInput{
				MaxNumberOfMessages: aws.Int64(1),
				QueueUrl:            aws.String(QueueURL),
				WaitTimeSeconds:     aws.Int64(20),
			}

			result, err := svc.ReceiveMessage(receiveParams)
			if err != nil {
				fmt.Println(err)
				time.Sleep(1 * time.Second)
				continue
			}

			for _, msg := range result.Messages {
				fmt.Println("Message received: %s \n", *msg.Body)

				deleteParams := &sqs.DeleteMessageInput{
					QueueUrl:      aws.String(QueueURL),
					ReceiptHandle: msg.ReceiptHandle,
				}

				_, err := svc.DeleteMessage(deleteParams)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
