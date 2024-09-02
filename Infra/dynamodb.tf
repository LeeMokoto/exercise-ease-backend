resource "aws_dynamodb_table" "user-table" {
  name           = "EE-User-Table"
  billing_mode   = "PAY_PER_REQUEST"
  range_key       = "UserId"
  hash_key      = "OrganisationID"
  

  attribute {
    name = "UserId"
    type = "S"
  }
  attribute {
    name = "OrganisationID"
    type = "S"
  }

  tags = {
    Name        = "ee-dynamodb-user-table"
    Environment = "production"
  }
}

resource "aws_dynamodb_table" "org-table" {
  name           = "EE-Org-Table"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "OrganisationID"
  range_key      = "OrganisationOwnerID"
  

  attribute {
    name = "OrganisationID"
    type = "S"
  }
  attribute {
    name = "OrganisationOwnerID"
    type = "S"
  }

  tags = {
    Name        = "ee-dynamodb-org-table"
    Environment = "production"
  }
}

resource "aws_dynamodb_table" "client-table" {
  name           = "EE-Client-Table"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "UserId"
    range_key      = "ClientId"
  

  attribute {
    name = "UserId"
    type = "S"
  }

    attribute {
    name = "ClientId"
    type = "S"
  }

  tags = {
    Name        = "ee-dynamodb-client-table"
    Environment = "production"
  }

  
}

resource "aws_dynamodb_table" "program-table" {
  name           = "EE-Programs-Table"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "ProgramId"
    range_key      = "ProgramName"
  

  attribute {
    name = "ProgramId"
    type = "S"
  }

    attribute {
    name = "ProgramName"
    type = "S"
  }

  tags = {
    Name        = "ee-dynamodb-program-table"
    Environment = "production"
  }

  
}

# resource "aws_dynamodb_table" "exercise-table" {
#   name           = "EE-Exercise-Table"
#   billing_mode   = "PAY_PER_REQUEST"
#   hash_key       = "ExerciseId"
#     range_key      = "ExerciseName"
  

#   attribute {
#     name = "ProgramId"
#     type = "S"
#   }

#     attribute {
#     name = "ProgramName"
#     type = "S"
#   }

#   tags = {
#     Name        = "ee-dynamodb-program-table"
#     Environment = "production"
#   }

  
# }