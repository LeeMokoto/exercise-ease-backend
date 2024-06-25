terraform {
  backend "s3" {
        bucket = "mojima-terraform-state"
        region = "eu-west-1"
        key = "exercise-ease/states/terraform.tfstate"
      
    }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.50.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.6.2"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "~> 2.2.0"
    }
  }

  required_version = "~>1.8.2"
}

provider "aws" {
  region  = "af-south-1"
}

