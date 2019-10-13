command build {
  expression = "make build"
}

command test {
  expression = <<EOF
    aws cloudformation deploy
      --template-file some-template-file
      --role-arn some-arn
      --stack-name ai-papi
      --parameter-overrides
          some_key=some_value
          some_key_some=some_value_some
  EOF
}

task also_test {
  script {
    concurrent = true

    run command test {}
    run command test {}
    run task another {}
  }
}
