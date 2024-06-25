
///get functions
variable "getFunction_names" {
  description = "Names of the getFunctions"
  type        = list(string)
  default     = ["getUser.zip"]
}

variable "getFunctions" {
  description = "Create get functions for database tables"
  type        = list(string)
  default     = ["getUser"]
}

variable "getFunctions_handlers" {
  description = "API handlers for getFunctions set"
  type        = list(string)
  default     = ["getUser"]
}

///create functions
variable "createFunction_names" {
  description = "Names of the createFunctions"
  type        = list(string)
  default     = ["createUser.zip"]
}

variable "createFunctions" {
  description = "Create get functions for database tables"
  type        = list(string)
  default     = ["createUser"]
}

variable "createFunctions_handlers" {
  description = "API handlers for getFunctions set"
  type        = list(string)
  default     = ["createUser"]
}