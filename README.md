---
marp: false
---

# aws-codecommit-pull-request-aws-codebuild-aws-chatbot
This repository contains sample code that tests pull requests created in AWS CodeCommit with CodeBuild and notifies to Slack using AWS Chatbot.

referenced in the qiita post:
XXXXXXXX

![Architecture](./doc/architecture.drawio.svg)


## Getting Started
### Prerequisites
 - macOS Catalina 10.15.6
 - go version go1.15.2 darwin/amd64
 - aws-cli/2.0.17
 - SAM CLI, version 1.2.0
 - Slack 4.8.0

Before you get started, make sure you've completed the tasks that Create new chatbot client.

<img width="60%" alt="chatbot_client_setting1" src="./doc/new_aws_chatbot_client.png">

<img width="60%" alt="hatbot_client_setting2" src="./doc/new_chatbot_client_slack.png">

<img width="60%" alt="hatbot_client_setting3" src="./doc/aws-chatbot-integration.png">

### Build
```
$ make
```

### Deployment
```
$ sam deploy --guided --capabilities CAPABILITY_NAMED_IAM
```


### Reference
[Getting started with AWS Chatbot](https://docs.aws.amazon.com/chatbot/latest/adminguide/getting-started.html)

[Validating AWS CodeCommit Pull Requests with AWS CodeBuild and AWS Lambda](https://aws.amazon.com/jp/blogs/devops/validating-aws-codecommit-pull-requests-with-aws-codebuild-and-aws-lambda/)

### Link
Qiita URL

