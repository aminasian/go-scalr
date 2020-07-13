resource "null_resource" "test" {
}

variable "testin" {
  default = "defaultval"
}

output "testout" {
  value = null_resource.test.id
}
