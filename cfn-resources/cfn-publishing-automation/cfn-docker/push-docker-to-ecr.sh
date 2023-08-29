#!/bin/bash
# This script is used to push the docker image to ECR
# It is used by the CodeBuild project

aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 711489243244.dkr.ecr.us-east-1.amazonaws.com

# Build the image
docker build -t atlas-cfn-publish-image .

# Tag the image
docker tag atlas-cfn-publish-image:latest 711489243244.dkr.ecr.us-east-1.amazonaws.com/atlas-cfn-publish-image:latest

# Push the image
docker push 711489243244.dkr.ecr.us-east-1.amazonaws.com/atlas-cfn-publish-image:latest
