import boto3


def lambda_handler(event, context):
    ecs_client = boto3.client('ecs')
    cluster_name = 'cluster'
    service_name = 'service'

    response = ecs_client.update_service(
        cluster=cluster_name,
        service=service_name,
        desiredCount=0
    )

    return {
        'statusCode': 200,
        'body': response
    }
