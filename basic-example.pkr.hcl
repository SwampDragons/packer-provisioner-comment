source "null" "basic-example" {
  ssh_host = "127.0.0.1"
  ssh_username = "foo"
  ssh_password = "bar"
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