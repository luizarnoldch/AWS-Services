sam package --template-file template.yml --output-template-file packaged.yml --s3-bucket amazon-connect-31732095774d


sam deploy --template-file packaged.yml --stack-name DynamoLambda --capabilities CAPABILITY_IAM

aws cloudformation delete-stack --stack-name DynamoLambda