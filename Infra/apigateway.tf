resource "aws_apigatewayv2_api" "main" {
  name        = "Exercise-Ease-API"
  description = "API for Exercise Ease"
  protocol_type = "HTTP"
  cors_configuration {
    allow_origins = ["*"]
    allow_methods = ["POST", "GET", "PUT", "OPTIONS", "DELETE"]
    allow_headers = ["content-type"]
    max_age = 300
  }
}

resource "aws_apigatewayv2_stage" "prod" {
  api_id = aws_apigatewayv2_api.main.id
  name        = "prod"
  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.main_api_gw.arn

    format = jsonencode({
      requestId               = "$context.requestId"
      sourceIp                = "$context.identity.sourceIp"
      requestTime             = "$context.requestTime"
      protocol                = "$context.protocol"
      httpMethod              = "$context.httpMethod"
      resourcePath            = "$context.resourcePath"
      routeKey                = "$context.routeKey"
      status                  = "$context.status"
      responseLength          = "$context.responseLength"
      integrationErrorMessage = "$context.integrationErrorMessage"
      }
    )
  }
}

resource "aws_cloudwatch_log_group" "main_api_gw" {
  name = "/aws/api-gw/${aws_apigatewayv2_api.main.name}"

  retention_in_days = 30
}

resource "aws_apigatewayv2_integration" "getFunctions_handler" {
  count = length(aws_lambda_function.getFunctions)
  api_id = aws_apigatewayv2_api.main.id

  integration_type = "AWS_PROXY"
  integration_uri  = aws_lambda_function.getFunctions[count.index].invoke_arn
}

resource "aws_apigatewayv2_integration" "createFunctions_handler" {
  count = length(aws_lambda_function.createFunctions)
  api_id = aws_apigatewayv2_api.main.id

  integration_type = "AWS_PROXY"
  integration_uri  = aws_lambda_function.createFunctions[count.index].invoke_arn
}

resource "aws_apigatewayv2_integration" "updateFunctions_handler" {
  count = length(aws_lambda_function.updateFunctions)
  api_id = aws_apigatewayv2_api.main.id

  integration_type = "AWS_PROXY"
  integration_uri  = aws_lambda_function.updateFunctions[count.index].invoke_arn
}

resource "aws_apigatewayv2_integration" "deleteFunctions_handler" {
  count = length(aws_lambda_function.deleteFunctions)
  api_id = aws_apigatewayv2_api.main.id

  integration_type = "AWS_PROXY"
  integration_uri  = aws_lambda_function.deleteFunctions[count.index].invoke_arn
}


# resource "aws_apigatewayv2_integration" "getPresignedURL_handler" {
#   api_id = aws_apigatewayv2_api.main.id

#   integration_type = "AWS_PROXY"
#   integration_uri  = aws_lambda_function.getPresignedURL.invoke_arn
# }

resource "aws_apigatewayv2_route" "getFunctions_handler" {
  count     = length(var.getFunctions)
  api_id    = aws_apigatewayv2_api.main.id
  route_key = "GET /${var.getFunctions[count.index]}"

  target = "integrations/${aws_apigatewayv2_integration.getFunctions_handler[count.index].id}"
}

resource "aws_apigatewayv2_route" "createFunctions_handler" {
  count     = length(var.createFunctions)
  api_id    = aws_apigatewayv2_api.main.id
  route_key = "POST /${var.createFunctions[count.index]}"

  target = "integrations/${aws_apigatewayv2_integration.createFunctions_handler[count.index].id}"
}

resource "aws_apigatewayv2_route" "updateFunctions_handler" {
  count     = length(var.updateFunctions)
  api_id    = aws_apigatewayv2_api.main.id
  route_key = "POST /${var.updateFunctions[count.index]}"

  target = "integrations/${aws_apigatewayv2_integration.updateFunctions_handler[count.index].id}"
}

resource "aws_apigatewayv2_route" "deleteFunctions_handler" {
  count     = length(var.deleteFunctions)
  api_id    = aws_apigatewayv2_api.main.id
  route_key = "DELETE /${var.deleteFunctions[count.index]}"

  target = "integrations/${aws_apigatewayv2_integration.deleteFunctions_handler[count.index].id}"
}

# resource "aws_apigatewayv2_route" "getPresignedURL_handler" {
#   api_id = aws_apigatewayv2_api.main.id
#   route_key = "GET /getPresignedURL"

#   target = "integrations/${aws_apigatewayv2_integration.getPresignedURL_handler.id}"
# }

resource "aws_lambda_permission" "getFunctions_api_gw" {
  count         = length(aws_lambda_function.getFunctions)
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.getFunctions[count.index].function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.main.execution_arn}/*/*"
}

resource "aws_lambda_permission" "createFunctions_api_gw" {
  count         = length(aws_lambda_function.createFunctions)
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.createFunctions[count.index].function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.main.execution_arn}/*/*"
}

resource "aws_lambda_permission" "updateFunctions_api_gw" {
  count         = length(aws_lambda_function.updateFunctions)
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.updateFunctions[count.index].function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.main.execution_arn}/*/*"
}

resource "aws_lambda_permission" "deleteFunctions_api_gw" {
  count         = length(aws_lambda_function.deleteFunctions)
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.deleteFunctions[count.index].function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.main.execution_arn}/*/*"
}

# resource "aws_lambda_permission" "getPresignedURL_api_gw" {
#   statement_id  = "AllowExecutionFromAPIGateway"
#   action        = "lambda:InvokeFunction"
#   function_name = aws_lambda_function.getPresignedURL.function_name
#   principal     = "apigateway.amazonaws.com"

#   source_arn = "${aws_apigatewayv2_api.main.execution_arn}/*/*"
# }

output "api_gateway_invoke_url" {
  description = "API gateway default stage invokation URL"
  value       = aws_apigatewayv2_stage.prod.invoke_url
}

