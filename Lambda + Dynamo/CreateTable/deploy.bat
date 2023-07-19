@echo off

set AWS_ACCESS_KEY_ID=AKIAZ5IN7MZSQWUHSK7A
set AWS_SECRET_ACCESS_KEY=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
set function_name=%1
if "%function_name%"=="" set function_name=myFunction
set runtime=go1.x
set handler=main
set role_arn=arn:aws:iam::681318508133:role/admin-test-role
set zip_file=main.zip
set region=%2
if "%region%"=="" set region=us-east-1

aws lambda create-function --function-name %function_name% --runtime %runtime% --handler %handler% --role %role_arn% --zip-file fileb://%zip_file% --region %region%

aws lambda create-function --function-name myFunction --runtime go1.x --handler main --role arn:aws:iam::681318508133:role/admin-test-role --zip-file fileb://main.zip --region us-east-1

echo Function created successfully.
