// Terraform configuration to use together with Localstack
// Resources created in this section will only be used to run the test suite.
// After starting the localstack, you can apply this Terraform config to create
// a fully functional testing environment. See README.md for more details.
provider "aws" {
  access_key                  = "mock_access_key"
  region                      = "us-east-1"
  s3_force_path_style         = true
  secret_key                  = "mock_secret_key"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    iam            = "http://localhost:4566"
    lambda         = "http://localhost:4566"
    s3             = "http://localhost:4566"
  }
}

resource "aws_s3_bucket" "test_bucket" {
  bucket        = "test-bucket"
  acl           = "public-read-write"
}

resource "aws_s3_bucket_object" "test_lambda" {
  bucket = aws_s3_bucket.test_bucket.bucket
  key    = "test-lambda/1.0.0/index.zip"
  source = "fixtures/index.zip"
  etag = filemd5("fixtures/index.zip")
}

data "aws_iam_policy_document" "test_lambda" {
  version = "2012-10-17"

  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"

    principals {
      type       = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "test_lambda" {
  name               = "test-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.test_lambda.json
  description        = "A role that can be assumed by the test lambda"
}

data "aws_iam_policy" "test_lambda_execution" {
  arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "test_lambda_execution" {
  role       = aws_iam_role.test_lambda.id
  policy_arn = data.aws_iam_policy.test_lambda_execution.arn
}


resource "aws_lambda_function" "test_lambda" {
  function_name = "test-lambda"
  description   = "A lambda function used for testing purposes"

  handler     = "main"
  runtime     = "go1.x"
  memory_size = 128
  timeout     = 5

  s3_bucket = aws_s3_bucket.test_bucket.bucket
  s3_key    = aws_s3_bucket_object.test_lambda.key

  role = aws_iam_role.test_lambda.arn
}

resource "aws_lambda_alias" "test_lambda" {
  name             = "live"
  description      = "Live version of the lambda. Gradual traffic shifting in place."
  function_name    = aws_lambda_function.test_lambda.arn
  function_version = "$LATEST"

  lifecycle {
    ignore_changes = [function_version]
  }
}
