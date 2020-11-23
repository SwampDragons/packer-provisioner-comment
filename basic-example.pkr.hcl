source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = ["sources.null.basic-example"]

  provisioner "comment" {
    comment = "Begin"
    ui = true
    bubble_text =  true
  }

  provisioner "shell-local"{
     inline = ["echo \"This is a shell script\""]
  }

  provisioner "comment" {
    comment = "In the middle of Provisioning run"
    ui = true
  }

  provisioner "shell-local"{
    inline = ["echo \"This is another shell script\""]
  }

  provisioner "comment" {
    comment = "this comment is invisible and won't go to the UI"
  }

  provisioner "comment" {
    comment = "End"
    ui = true
    bubble_text =  true
  }
}