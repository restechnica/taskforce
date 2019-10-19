command build {
  text = "go build -o 'bin/taskforce' github.com/restechnica/taskforce/cmd/taskforce"
}

command build {
  text = "go build -o 'bin/taskforce' github.com/restechnica/taskforce/cmd/taskforce"
}

command clean {
  text = "rm -rf bin"
}

command install {
  text = "go install github.com/restechnica/taskforce/cmd/taskforce"
}

command test {
  text = "echo ${script("python", "vars.py", "some_var.hello")} ${script("python", "./vars.py", "another_var")} ${script("python", "./vars.py", "another_var")}"
}

task test {
  run command test {}
}
