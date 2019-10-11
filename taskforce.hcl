command build {
  expression = "go build -o 'bin/taskforce' github.com/restechnica/taskforce/cmd/taskforce"
}

command clean {
  expression = "rm -rf bin"
}

command install {
  expression = "go install github.com/restechnica/taskforce/cmd/taskforce"
}

command test {
  expression = <<EOF
    aws cloudformation deploy
      --template-file some-template-file
      --role-arn some-arn
      --stack-name ai-papi
      --parameter-overrides some_key=some_value some_key_some=some_value_some
  EOF
}
