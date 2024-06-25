
///get functions
variable "getFunction_names" {
  description = "Names of the getFunctions"
  type        = list(string)
  default     = ["getUser.zip","getClient.zip","getOrg.zip"]
}

variable "getFunctions" {
  description = "Create get functions for database tables"
  type        = list(string)
  default     = ["getUser","getClient","getOrg"]
}

variable "getFunctions_handlers" {
  description = "API handlers for getFunctions set"
  type        = list(string)
  default     = ["getUser","getClient","getOrg"]
}

variable "getFunction_path"{
  description = "path of code"
  type = list(string)
  default = ["getUser/bootstrap.zip","getClient/bootstrap.zip","getOrg/bootstrap.zip"]
}

///create functions
variable "createFunction_names" {
  description = "Names of the createFunctions"
  type        = list(string)
  default     = ["createUser.zip","createClient.zip","createOrg.zip"]
}

variable "createFunctions" {
  description = "Create get functions for database tables"
  type        = list(string)
  default     = ["createUser","createClient","createOrg"]
}

variable "createFunctions_handlers" {
  description = "API handlers for getFunctions set"
  type        = list(string)
  default     = ["createUser","createClient","createOrg"]
}

variable "createFunction_path"{
  description = "path of code"
  type = list(string)
  default = ["createUser/bootstrap.zip","createClient/bootstrap.zip","createOrg/bootstrap.zip"]
}