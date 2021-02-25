packer {
  required_plugins {
    comment = {
      version = ">=v0.2.23"
      source  = "github.com/sylviamoss/comment"
    }
  }
}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = ["sources.null.basic-example"]

  provisioner "comment" {
    comment = "Basic example to test with latest packer"
    ui = true
  }
}