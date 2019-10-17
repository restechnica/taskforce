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
  text = "echo test"
}

task test {
  run command test {}
}
