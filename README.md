# Taskforce

Taskforce is a tool for defining, running and versioning operational tasks and commands. Taskforce can be used in conjunction with the platform and technologies of your choice. It is available for Linux, macOS and Windows.

The key features of Taskforce are:

- **Operations as Code**: Operations are described using a high-level configuration syntax. This allows a blueprint of your operations to be versioned and treated as you would any other code. Additionally operations can be shared and re-used by both man and machine.

- **Environment Interpolation**: Taskforce comes with a `dotenv` implementation out of the box. This allows you to set up a .env file with environment variables and work with those variables within tasks and commands. It is an abstraction which allows different environments, developer or pipeline, to run operational tasks in identical fashion yet with enviromental differences.

- **Script Interpolation**: Setting environment variables does not always cut it. Sometimes you just need more complex logic or HTTP calls to get the information you require in your operations, especially in the age of Cloud. Taskforce allows you to use variables from scripts in your language(s) of choice within its configuration.

- **Cross-platform**: All major platforms are supported.

Taskforce has taken parts of the feature set and ideas from other tools like [Serverless.com](https://github.com/serverless/serverless) and [Terraform](https://github.com/hashicorp/terraform) but has made them available outside of those tools with the **technology of your choice**.

## License
[Mozilla Public License v2.0](https://github.com/hashicorp/terraform/blob/master/LICENSE)



