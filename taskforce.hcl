command build {
  expression = "go build -o 'bin/taskforce' github.com/restechnica/taskforce/cmd/taskforce"
}

command clean {
  expression = "rm -rf bin"
}

command install {
  expression = "go install github.com/restechnica/taskforce/cmd/taskforce"
}
