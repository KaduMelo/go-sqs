# SQS in GO

### Tools already installed

1. Docker
2. AWS Cli
3. Localstack

### Comands to be utilized:

```docker run --rm -it -d -p 4566:4566 localstack/localstack```

```aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name my-first-queue```

```aws --endpoint-url=http://localhost:4566 sqs send-message --queue-url http://sqs.us-west-2.localhost.localstack.cloud:4566/000000000000/my-first-queue --message-body "Message of the test"```