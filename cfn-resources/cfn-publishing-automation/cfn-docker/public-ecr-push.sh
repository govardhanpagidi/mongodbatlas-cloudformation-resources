#!/bin/bash
# This script is used to push the docker image to ECR
# It is used by the CodeBuild project

#aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 711489243244.dkr.ecr.us-east-1.amazonaws.com

#aws_regions="us-east-1"
#aws_regions="us-east-1 us-east-2 us-west-1 us-west-2 eu-west-1 eu-central-1 ap-southeast-1 ap-southeast-2 ap-northeast-1 ap-northeast-2 ap-south-1 sa-east-1 ca-central-1 eu-west-2 eu-west-3 eu-north-1 eu-south-1 me-south-1 af-south-1"
#for region in $aws_regions; do
   aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/k5o3c0z4

    # Build the image
    docker build -t atlas-cfn-public-image .

    # Tag the image
   docker tag atlas-cfn-public-image:latest public.ecr.aws/k5o3c0z4/atlas-cfn-public-image:latest
     # Push the image
   docker push public.ecr.aws/k5o3c0z4/atlas-cfn-public-image:latest
#done
