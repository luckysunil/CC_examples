variable "region" {
  type = string
  description = "AWS Deployment Region"
  default = "ap-south-1"
}

variable "cidr_block" {
  type = string
  description = "VPC CIDR Block"
  default = "10.10.0.0/16"
}

variable "public_subnet_cidrs" {
 type        = list(string)
 description = "Public Subnet CIDR values"
 default     = ["10.10.1.0/24", "10.10.2.0/24", "10.10.3.0/24"]
}

variable "private_subnet_cidrs" {
 type        = list(string)
 description = "Private Subnet CIDR values"
 default     = ["10.10.4.0/24", "10.10.5.0/24", "10.10.6.0/24"]
}

variable "azs" {
 type        = list(string)
 description = "Availability Zones"
 default     = ["ap-south-1a", "ap-south-1b", "ap-south-1c"]
}
