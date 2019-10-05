resource "task" simple {
  command = ""

  variables {
    some_local_overwrite = ""
  }
}

resource "pipeline" some_pipeline_name {
  branches = [
  ]
  variables {
  }

  step {
    artifacts = [
    ]
    caches = [
    ]
    tasks = [
    ]
    trigger = "manual"
    variables {
    }
  }
}
