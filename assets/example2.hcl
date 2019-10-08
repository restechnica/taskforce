task "clean:build" {
  commands = [
    series([
      "clean",
      "build"
    ])
  ]
}

task "crazy" {

  script {
    actions = [
      task("<name>"),
      command("clean"),
      command("build"),

      parallel([
        command("run"),
        command("install")
      ])
    ]
  }
}


/**
best one for now
*/
task "ex" {
  script {
    action task "crazy" {}
    action command "clean" {}
    action command "build" {}
  }

  script "inline" {
    parallel = true

    action command "run" {}
    action command "install" {}
  }

  script "inline" {
    action script "miight-be-cool" {}
    action command "install" {}
  }

  script {
    file = file('../some/example')
  }
}
