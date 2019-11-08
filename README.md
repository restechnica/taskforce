# Taskforce

Taskforce is a tool for defining, running and versioning operational tasks and commands.

The key features of Taskforce are:

- **Operations as Code**: Operations are described using a high-level configuration syntax. This allows a blueprint of your operations to be versioned and treated as you would any other code. Additionally operations can be shared and re-used by both man and machine.

- **Environment Interpolation**: Taskforce comes with a `dotenv` implementation out of the box. This allows you to set up a .env file with environment variables and work with those variables within tasks and commands. It is an abstraction which allows different environments, developer or pipeline, to run operational tasks in identical fashion yet with enviromental differences.

- **Script Interpolation**: Setting environment variables does not always cut it. Sometimes you just need more complex logic or HTTP calls to get the information you require in your operations, especially in the age of Cloud. Taskforce allows you to use variables from scripts in your language(s) of choice within its configuration.

- **Cross-platform**: It is available for Linux, macOS and Windows.

Taskforce has taken parts of the feature set and ideas from other tools like [Serverless.com](https://github.com/serverless/serverless) and [Terraform](https://github.com/hashicorp/terraform) but has made them available outside of those tools with the **technology of your choice**.

## Getting Started

1. Create a `taskforce.hcl` file. It is usually located in the root of a project or repository.
Check out the [HCL GitHub page](https://github.com/hashicorp/hcl) for more information on HCL syntax.
2. Add tasks and commands. See the Configuration Reference below.
3. Run those tasks and commands with the Taskforce CLI.

## Command

Defines a command and makes it available through the CLI.

### Usage
```shell script
$ taskforce -c <name>
```

### Configuration Reference
#### block
- `command`

#### labels
- `Name` string - (Required) The name of the command. Referenced by tasks and the CLI.

#### attributes
- `Directory` string - (Optional) The working directory to run the command in.
Relative and absolute paths are supported.Tilde `~` is  supported.

- `Text` string - (Required) The command text. Multiline supported.

### Example
```hcl
command example {
  directory = "~"
  text = "ls -a"
}
```
```shell script
$ taskforce -c example
```

#### multiline
```hcl
command multiline_example {
  directory = "~"
  text = <<EOF
    aws cloudformation deploy
        --capabilities CAPABILITY_IAM
        --role-arn some_role
        --stack-name some_stack
        --template-file some_file
  EOF
}
```
```shell script
$ taskforce -c multiline_example
```

## Task

A task has instructions. Instructions can be used to run commands and other tasks.
The instructions are executed sequentially.

### Usage
```shell script
$ taskforce <name>
```

### Configuration Reference
#### block
- `task`

#### labels
- `Name` string - (Required) The name of the task. Referenced by other tasks and the CLI.

#### nested blocks
- `run` Instruction
    - `Type` *label,string* - (Required) Type of the instruction. can be either `task` or `command`.
    - `Reference` *label,string* - (Required) Name of the task or command.

### Example
```hcl
task example {
  run command some_command {}
  run task some_task {}
}
```

```shell script
$ taskforce example
```

## Environment Variables
With Taskforce you have your environment variables at your disposal. You can reference them in the HCL configuration
and they are passed on to the commands that Taskforce executes.
Each environment variable is available in upper and lowercase.

### Usage
```hcl
command env_var_example {
  directory = "~"
  text = <<EOF
    aws cloudformation deploy
        --capabilities CAPABILITY_IAM
        --role-arn ${env.role_arn}      // or ${env.ROLE_ARN}
        --stack-name ${env.stack_name}  // or ${env.STACK_NAME}
        --template-file some_file
  EOF
}
```

### Custom environment variables
Taskforce supports the definition of environment variables.
Each command will be able to reference these environment variables as shown in the 'Usage'
section above. See the [dotenv project](https://github.com/bkeepers/dotenv) and its [Golang implementation](https://github.com/joho/godotenv) on GitHub for more information.

#### Variables with values not shared between environments

The variables are defined in a `.env` file in the root of the project. This file should __not__ be under version control and should only be used for development purposes. 
1. Create a `.env` file in the same directory as your Taskforce configuration file.
2. Define environment variables as key-value pairs with a `=` inbetween keys and values.
3. Reference the environment variables in the Taskforce configuration file as you would any other environment variable.

A .env file when working with AWS and CloudFormation could look like the following:
```
aws_access_key_id=AKIA2UI5876HJJ4WQH4
aws_secret_access_key=mrWMk01234Gct9876v6IDSpJ8yLpfDjLtwF
aws_region=eu-west-1
role_arn=arn:...:...:role:...
stack_name=stack-name
```

#### Variables with values shared between environments

The variables are defined in a `taskforce.env` file located in the project root.
This file is meant to be under version control and should define variables which hold the same value between environments.

The key value pair syntax is exactly the same as the '.env' syntax shown above.

## JSON variables

Taskforce supports JSON variable references.

### Example

```hcl
command json_var_example {
  text = <<EOF
    aws cloudformation deploy
        --capabilities CAPABILITY_IAM
        --region ${json("test.json", "aws.region")}
        --stack-name ${json("test.json", "aws.stack.name")}
        --template-file some_file
  EOF
}
```

## Script Variables

Taskforce supports script variable references for complex variable resolution. It is required that the script outputs the variables as JSON.
You can use a language of your choice and spread variables across different scripts.

### Script Example

#### Python
```python
#!/usr/bin/env python

import json

variables = {
    'some_var': {
        'a_nested_var': 'value'
    },
    'another_var': 'value'
}

var_json = json.dumps(variables)
print(var_json)
```
```hcl
command script_var_example {
  directory = "~"
  text = <<EOF
    aws cloudformation deploy
        --capabilities CAPABILITY_IAM
        --role-arn ${script("python", "variables.py", "another_var")} 
        --stack-name ${script("python", "variables.py", "some_var.a_nested_var")}
        --template-file some_file
  EOF
}
```

#### Node.js
```javascript
#!/usr/bin/env node

const variables = {
    some_var: {
        a_nested_var: 'value'
    },
    another_var: 'value'
};

const json = JSON.stringify(variables);

console.log(json);
```
```hcl
command script_var_example {
  directory = "~"
  text = <<EOF
    aws cloudformation deploy
        --capabilities CAPABILITY_IAM
        --role-arn ${script("node", "variables.js", "another_var")} 
        --stack-name ${script("node", "variables.js", "some_var.a_nested_var")}
        --template-file some_file
  EOF
}
```


## License
[Mozilla Public License v2.0](./LICENSE)



