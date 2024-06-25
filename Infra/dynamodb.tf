resource "aws_dynamodb_table" "basic-dynamodb-table" {
  name           = "EE-User-Table"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "UserId"
  range_key      = "OrganisationID"
  

  attribute {
    name = "UserId"
    type = "S"
  }

  # attribute {
  #   name = "Name"
  #   type = "S"
  # }

  # attribute {
  #   name = "Surname"
  #   type = "S"
  # }
  # attribute {
  #   name = "Email"
  #   type = "S"
  # }
  attribute {
    name = "OrganisationID"
    type = "S"
  }

  ttl {
    attribute_name = "TimeToExist"
    enabled        = true
  }

#   global_secondary_index {
#     name               = "GameTitleIndex"
#     hash_key           = "GameTitle"
#     range_key          = "TopScore"
#     write_capacity     = 10
#     read_capacity      = 10
#     projection_type    = "INCLUDE"
#     non_key_attributes = ["UserId"]
#   }

  tags = {
    Name        = "ee-dynamodb-user-table"
    Environment = "production"
  }
}