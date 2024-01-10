package main

import (
	"context"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func handler(ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}
	clientEcs := ecs.NewFromConfig(cfg)
	serviceArn := "arn:aws:ecs:us-east-1:*:service/cluster/service"
	_, err = clientEcs.UpdateService(ctx, &ecs.UpdateServiceInput{
		Service:      aws.String(serviceArn),
		DesiredCount: aws.Int32(0),
	})
	if err != nil {
		return err
	}
	slog.Info("Serviço ECS atualizado para 0 tarefas")
	return nil
}

func main() {
	lambda.Start(handler)
}
