
///get functions
variable "getFunction_names" {
  description = "Names of the getFunctions"
  type        = list(string)
  default     = ["getUser.zip","getClients.zip","getOrg.zip","getPrograms.zip","getClientProfile.zip","getUsers.zip"]
}

variable "getFunctions" {
  description = "Create get functions for database tables"
  type        = list(string)
  default     = ["getUser","getClients","getOrg","getPrograms","getClientProfile","getUsers"]
}

variable "getFunctions_handlers" {
  description = "API handlers for getFunctions set"
  type        = list(string)
  default     = ["getUser","getClients","getOrg","getPrograms","getClientProfile","getUsers"]
}

variable "getFunction_path"{
  description = "path of code"
  type = list(string)
  default = ["getUser/bootstrap.zip","getClients/bootstrap.zip","getOrg/bootstrap.zip","getPrograms/bootstrap.zip","getClientProfile/bootstrap.zip","getUsers/bootstrap.zip"]
}

///update functions
variable "updateFunction_names" {
  description = "Names of the updateFunctions"
  type        = list(string)
  default     = ["updateUser.zip","updateClient.zip","updateOrg.zip"]
}

variable "updateFunctions" {
  description = "Create get functions for database tables"
  type        = list(string)
  default     = ["updateUser","updateClient","updateOrg"]
}

variable "updateFunctions_handlers" {
  description = "API handlers for updateFunctions set"
  type        = list(string)
  default     = ["updateUser","updateClient","updateOrg"]
}

variable "updateFunction_path"{
  description = "path of code"
  type = list(string)
  default = ["updateUser/bootstrap.zip","updateClient/bootstrap.zip","updateOrg/bootstrap.zip"]
}

///delete functions
variable "deleteFunction_names" {
  description = "Names of the deleteFunctions"
  type        = list(string)
  default     = ["deleteUser.zip","deleteClient.zip","deleteProgram.zip"]
}

variable "deleteFunctions" {
  description = "Create get functions for database tables"
  type        = list(string)
  default     = ["deleteUser","deleteClient","deleteProgram"]
}

variable "deleteFunctions_handlers" {
  description = "API handlers for deleteFunctions set"
  type        = list(string)
  default     = ["deleteUser","deleteClient","deleteProgram"]
}

variable "deleteFunction_path"{
  description = "path of code"
  type = list(string)
  default = ["deleteUser/bootstrap.zip","deleteClient/bootstrap.zip","deleteProgram/bootstrap.zip"]
}

///create functions
variable "createFunction_names" {
  description = "Names of the createFunctions"
  type        = list(string)
  default     = ["createUser.zip","createClient.zip","createOrg.zip", "adminCreateUser.zip","createProgram.zip"]
}

variable "createFunctions" {
  description = "Create get functions for database tables"
  type        = list(string)
  default     = ["createUser","createClient","createOrg","adminCreateUser","createProgram"]
}

variable "createFunctions_handlers" {
  description = "API handlers for getFunctions set"
  type        = list(string)
  default     = ["createUser","createClient","createOrg", "adminCreateUser","createProgram"]
}

variable "createFunction_path"{
  description = "path of code"
  type = list(string)
  default = ["createUser/bootstrap.zip","createClient/bootstrap.zip","createOrg/bootstrap.zip", "adminCreateUser/bootstrap.zip", "createProgram/bootstrap.zip"]
}


//default variables
variable "region" {
  description = "The default AWS region to use for provisioning infrastructure"
  type        = string
  default     = "af-south-1"

  validation {
    condition     = can(regex("[a-z][a-z]-[a-z]+-[1-9]", var.region))
    error_message = "Must be valid AWS region name"
  }
}

variable "project_name" {
  description = "The name of the project used for tagging resources"
  type        = string
  default     = "exercise-ease"
}

//vpc 
variable "vpc_data" {
  description = "Variable to hold and object for vpc configuration"

  type = object({
    vpc_cidr = string
    availability_zones = list(object({
      az_name                 = string
      public_subnet_cidr      = string
      private_app_subnet_cidr = string
      private_db_subnet_cidr  = string
    }))
  })

  default = {
    vpc_cidr = "10.0.0.0/16"
    availability_zones = [{
      az_name                 = "af-south-1a"
      public_subnet_cidr      = "10.0.0.0/20"
      private_app_subnet_cidr = "10.0.64.0/20"
      private_db_subnet_cidr  = "10.0.128.0/20"
      },
      {
        az_name                 = "af-south-1b"
        public_subnet_cidr      = "10.0.16.0/20"
        private_app_subnet_cidr = "10.0.80.0/20"
        private_db_subnet_cidr  = "10.0.144.0/20"
      },
      {
        az_name                 = "af-south-1c"
        public_subnet_cidr      = "10.0.32.0/20"
        private_app_subnet_cidr = "10.0.96.0/20"
        private_db_subnet_cidr  = "10.0.160.0/20"
    }]
  }
}









