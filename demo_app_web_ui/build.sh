#!/bin/bash -eu

echo "🔨 Build binary"
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "🎒 Create Zip file"
zip deployment.zip main

echo "🛀 Clean up"
rm main

echo "🔩 Create function on AWS"
aws lambda create-function --function-name $2 \
--zip-file fileb://./deployment.zip \
--runtime go1.x --handler main \
--role arn:aws:iam::$1:role/$3 \
--region $4

# echo "⚡️ Upload to AWS"
# aws lambda update-function-code --function-name $2 \
# --zip-file fileb://./deployment.zip \
# --region $4
