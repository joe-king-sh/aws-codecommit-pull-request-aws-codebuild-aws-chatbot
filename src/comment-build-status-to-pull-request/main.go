package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codecommit"
)

type EventDetail struct {
	SourceReference      string   `json:"sourceReference"`
	LastModifiedDate     string   `json:"lastModifiedDate"`
	Author               string   `json:"author"`
	PullRequestStatus    string   `json:"pullRequestStatus"`
	IsMerged             string   `json:"isMerged"`
	NotificationBody     string   `json:"notificationBody"`
	DestinationReference string   `json:"destinationReference"`
	PullRequestId        string   `json:"pullRequestId"`
	CallerUserArn        string   `json:"callerUserArn"`
	Title                string   `json:"title"`
	CreationDate         string   `json:"creationDate"`
	RepositoryNames      []string `json:"repositoryNames"`
	DestinationCommit    string   `json:"destinationCommit"`
	Event                string   `json:"event"`
	SourceCommit         string   `json:"sourceCommit"`
}

const region = "ap-northeast-1"

func HandleRequest(ctx context.Context, event events.CloudWatchEvent) (string, error) {

	mySession := session.Must(session.NewSession())

	// Get PR informations from CloudWatchEvents
	var eventDetail EventDetail
	err := json.Unmarshal(event.Detail, &eventDetail)
	if err != nil {
		print(err.Error())
		return "", err
	}
	pullRequestId := eventDetail.PullRequestId
	repositoryName := eventDetail.RepositoryNames[0]
	afterCommitId := eventDetail.DestinationCommit
	beforeCommitId := eventDetail.SourceCommit
	log.Printf("Target repository name is %s\n", repositoryName)
	log.Printf("PullRequests title is %s\n", eventDetail.Title)

	// Get CodeBuild badge url.
	codebuildSvc := codebuild.New(mySession, aws.NewConfig().WithRegion(region))
	codeBuildArn := os.Getenv("CODEBUILD_ARN")
	names := []*string{&codeBuildArn}
	batchGetProjectsOutput, err := codebuildSvc.BatchGetProjects(&codebuild.BatchGetProjectsInput{
		Names: names,
	})
	if err != nil {
		print(err.Error())
		return "", err
	}
	badgeUrl := batchGetProjectsOutput.Projects[0].Badge.BadgeRequestUrl

	// Set content to post to PR's comment
	commentTemplate := `Unit tests have been started in CodeBuild.  
Build Status:
![BuildBadge](%s)
`
	//Branch in badge URL is master by default, so, replace to 'develop branch'
	content := fmt.Sprintf(commentTemplate, strings.Replace(*badgeUrl, "master", "develop", -1))

	// Post comment for PR
	codecommitSvc := codecommit.New(mySession, aws.NewConfig().WithRegion(region))
	output, err := codecommitSvc.PostCommentForPullRequest(
		&codecommit.PostCommentForPullRequestInput{
			RepositoryName: &repositoryName,
			AfterCommitId:  &afterCommitId,
			BeforeCommitId: &beforeCommitId,
			Content:        &content,
			PullRequestId:  &pullRequestId,
		})
	if err != nil {
		print(err.Error())
		return "", err
	}

	log.Printf("Result for post comment for PullRequests:  %s\n", output)

	return output.GoString(), nil
}

func main() {
	lambda.Start(HandleRequest)
}
