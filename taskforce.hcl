command build {
  expression = "go build -o 'bin/taskforce' github.com/restechnica/taskforce/cmd/taskforce"
}

command clean {
  expression = "rm -rf bin"
}

command install {
  expression = "go install github.com/restechnica/taskforce/cmd/taskforce"
}

command run {
  directory = "~/Workspace/taskforce"
  expression = "echo ${env.AWS_SECRET_ACCESS_KEY}"
}

