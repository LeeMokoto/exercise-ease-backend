

resource "aws_lambda_function" "getFunctions"{
  count = length(var.getFunctions)
  function_name = var.getFunctions[count.index]
  timeout = 60
  # The bucket name as created earlier with "aws s3api create-bucket"
  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_object.getFunctions[count.index].key
  memory_size = 512
  handler = "bootstrap"
  runtime = "provided.al2023"
  
  source_code_hash = data.archive_file.getFunctions[count.index].output_base64sha256

  role = aws_iam_role.lambda_exec.arn
}

resource "aws_lambda_function" "createFunctions"{
  count = length(var.createFunctions)
  function_name = var.createFunctions[count.index]
  timeout = 60
  # The bucket name as created earlier with "aws s3api create-bucket"
  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_object.createFunctions[count.index].key
  memory_size = 512
  handler = "bootstrap"
  runtime = "provided.al2023"
  
  source_code_hash = data.archive_file.createFunctions[count.index].output_base64sha256

  role = aws_iam_role.lambda_exec.arn
}


# IAM role which dictates what other AWS services the Lambda function
# may access.
resource "aws_iam_role" "lambda_exec" {
  name = "rag_lambda_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"

}

data "aws_iam_policy_document" "policy" {
  statement {
    effect = "Allow"
    actions = ["dynamodb:*", "s3:*", "secretsmanager:*"]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "policy" {
  name        = "exercise-ease-policy"
  description = "Exercise Ease policy"
  policy      = data.aws_iam_policy_document.policy.json
}

resource "aws_iam_role_policy_attachment" "test-attach" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.policy.arn
}

resource "aws_cloudwatch_log_group" "getFunctions_lambda" {
  count = length(var.getFunctions)
  name = "/aws/lambda/${aws_lambda_function.getFunctions[count.index].function_name}"
}

resource "aws_cloudwatch_log_group" "createFunctions_lambda" {
  count = length(var.createFunctions)
  name = "/aws/lambda/${aws_lambda_function.createFunctions[count.index].function_name}"
}

//data for the lambda zip
data "archive_file" "getFunctions" {
  count       = length(var.getFunctions)
  type        = "zip"
  source_file  = "${path.module}/${var.getFunctions[count.index]}/bootstrap"
  output_path = "${path.module}/${var.getFunction_path[count.index]}"
}

data "archive_file" "createFunctions" {
  count       = length(var.createFunctions)
  type        = "zip"
  source_file  = "${path.module}/${var.createFunctions[count.index]}/bootstrap"
  output_path = "${path.module}/${var.createFunction_path[count.index]}"
}

resource "aws_s3_object" "getFunctions" {
  count = length(data.archive_file.getFunctions)
  bucket = aws_s3_bucket.lambda_bucket.id
  key    = "${var.getFunctions[count.index]}/bootstrap.zip"
  source = data.archive_file.getFunctions[count.index].output_path
  etag   = filemd5(data.archive_file.getFunctions[count.index].output_path)
}

resource "aws_s3_object" "createFunctions" {
  count = length(data.archive_file.createFunctions)
  bucket = aws_s3_bucket.lambda_bucket.id
  key    = "${var.createFunctions[count.index]}bootstrap.zip"
  source = data.archive_file.createFunctions[count.index].output_path
  etag   = filemd5(data.archive_file.createFunctions[count.index].output_path)
}

