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
 - SAM CLI, version 1.2.0
 - Slack 4.8.0

### Installing
なければとばす

### Deployment
#### 手動でやるchatbotとか
#### sam deployで方をつけたい
```
sam deploy -t template.yml --guided
```
最初は失敗する reason for CAPABILITY_NAMED_IAM 

```bash
sam deploy -t template.yml --capabilities CAPABILITY_NAMED_IAM

sam deploy -t template.yml --capabilities CAPABILITY_NAMED_IAM \
--parameter-overrides TargetWorkspaceId=T01B0FE4QM8 TargetChannelId=C01B658TYTZ
```

GOOS=linux go build main.go makeでこいつらできるように


### Reference

[Validating AWS CodeCommit Pull Requests with AWS CodeBuild and AWS Lambda](https://aws.amazon.com/jp/blogs/devops/validating-aws-codecommit-pull-requests-with-aws-codebuild-and-aws-lambda/)

### Link
Qiita URL

