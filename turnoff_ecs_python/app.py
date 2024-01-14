import boto3
from logging import info


def lambda_handler(event, context):
    ecs_client = boto3.client('ecs')
    service_arn = 'service_arn'

    response = ecs_client.update_service(
        service=service_arn,
        desiredCount=0
    )

    info(f'O serviço {service_arn} foi desligado!')

    return {
        'statusCode': 200,
        'body': response
    }
