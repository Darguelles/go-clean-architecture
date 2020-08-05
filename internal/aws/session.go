package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
)
// GetSession creates a session with AWS
func GetSession() *session.Session {
	sessConfig := aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID),
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            sessConfig,
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}
