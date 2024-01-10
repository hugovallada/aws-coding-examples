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
	clusterName := "hugovallada"

	_, err = clientEcs.DeleteCluster(ctx, &ecs.DeleteClusterInput{
		Cluster: aws.String(clusterName),
	})

	if err != nil {
		return err
	}

	slog.Info("ECS Cluster", clusterName, "was successfully deleted.")

	return nil
}

func main() {
	lambda.Start(handler)
}
