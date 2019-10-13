task some_task {
  command = ""

  variables {
    some_local_overwrite = ""
  }
}

pipeline some_pipeline_name {
  branches = [
  ]
  variables {
  }

  step {//runs tasks in parallel automatically, no sync options, new step needed for sync, makes more sense to allow tasks to run parallel
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
