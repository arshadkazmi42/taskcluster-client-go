// The following code is AUTO-GENERATED. Please DO NOT edit.
// To update this generated code, run the following command:
// in the /codegenerator/model subdirectory of this project,
// making sure that `${GOPATH}/bin` is in your `PATH`:
//
// go install && go generate
//
// This package was generated from the schema defined at
// https://references.taskcluster.net/github/v1/api.json

// The github service, typically available at
// `github.taskcluster.net`, is responsible for publishing pulse
// messages in response to GitHub events.
//
// This document describes the API end-point for consuming GitHub
// web hooks, as well as some useful consumer APIs.
//
// When Github forbids an action, this service returns an HTTP 403
// with code ForbiddenByGithub.
//
// See: https://docs.taskcluster.net/reference/core/github/api-docs
//
// How to use this package
//
// First create a Github object:
//
//  github := tcgithub.New(nil)
//
// and then call one or more of github's methods, e.g.:
//
//  err := github.Ping(.....)
//
// handling any errors...
//
//  if err != nil {
//  	// handle error...
//  }
//
// Taskcluster Schema
//
// The source code of this go package was auto-generated from the API definition at
// https://references.taskcluster.net/github/v1/api.json together with the input and output schemas it references, downloaded on
// Tue, 18 Sep 2018 at 16:23:00 UTC. The code was generated
// by https://github.com/taskcluster/taskcluster-client-go/blob/master/build.sh.
package tcgithub

import (
	"net/url"

	tcclient "github.com/taskcluster/taskcluster-client-go"
)

type Github tcclient.Client

// New returns a Github client, configured to run against production. Pass in
// nil to create a client without authentication. The
// returned client is mutable, so returned settings can be altered.
//
//  github := tcgithub.New(nil)                              // client without authentication
//  github.BaseURL = "http://localhost:1234/api/Github/v1"   // alternative API endpoint (production by default)
//  err := github.Ping(.....)                                // for example, call the Ping(.....) API endpoint (described further down)...
//  if err != nil {
//          // handle errors...
//  }
func New(credentials *tcclient.Credentials) *Github {
	return &Github{
		Credentials: credentials,
		ServiceName: "github",
		APIVersion:  "v1",
	}
}

// NewFromEnv returns a Github client with credentials taken from the environment variables:
//
//  TASKCLUSTER_CLIENT_ID
//  TASKCLUSTER_ACCESS_TOKEN
//  TASKCLUSTER_CERTIFICATE
//  TASKCLUSTER_ROOT_URL
//
// No validation is performed on the loaded values, and unset environment
// variables will result in empty string values.
//
// If environment variable TASKCLUSTER_CLIENT_ID is empty string or not set,
// authentication will be disabled.
func NewFromEnv() *Github {
	return &Github{
		Credentials: tcclient.CredentialsFromEnvVars(),
		ServiceName: "github",
		APIVersion:  "v1",
	}
}

// Respond without doing anything.
// This endpoint is used to check that the service is up.
//
// See https://docs.taskcluster.net/reference/core/github/api-docs#ping
func (github *Github) Ping() error {
	cd := tcclient.Client(*github)
	_, _, err := (&cd).APICall(nil, "GET", "/ping", nil, nil)
	return err
}

// Stability: *** EXPERIMENTAL ***
//
// Capture a GitHub event and publish it via pulse, if it's a push,
// release or pull request.
//
// See https://docs.taskcluster.net/reference/core/github/api-docs#githubWebHookConsumer
func (github *Github) GithubWebHookConsumer() error {
	cd := tcclient.Client(*github)
	_, _, err := (&cd).APICall(nil, "POST", "/github", nil, nil)
	return err
}

// Stability: *** EXPERIMENTAL ***
//
// A paginated list of builds that have been run in
// Taskcluster. Can be filtered on various git-specific
// fields.
//
// See https://docs.taskcluster.net/reference/core/github/api-docs#builds
func (github *Github) Builds(continuationToken, limit, organization, repository, sha string) (*BuildsResponse, error) {
	v := url.Values{}
	if continuationToken != "" {
		v.Add("continuationToken", continuationToken)
	}
	if limit != "" {
		v.Add("limit", limit)
	}
	if organization != "" {
		v.Add("organization", organization)
	}
	if repository != "" {
		v.Add("repository", repository)
	}
	if sha != "" {
		v.Add("sha", sha)
	}
	cd := tcclient.Client(*github)
	responseObject, _, err := (&cd).APICall(nil, "GET", "/builds", new(BuildsResponse), v)
	return responseObject.(*BuildsResponse), err
}

// Stability: *** EXPERIMENTAL ***
//
// Checks the status of the latest build of a given branch
// and returns corresponding badge svg.
//
// See https://docs.taskcluster.net/reference/core/github/api-docs#badge
func (github *Github) Badge(owner, repo, branch string) error {
	cd := tcclient.Client(*github)
	_, _, err := (&cd).APICall(nil, "GET", "/repository/"+url.QueryEscape(owner)+"/"+url.QueryEscape(repo)+"/"+url.QueryEscape(branch)+"/badge.svg", nil, nil)
	return err
}

// Stability: *** EXPERIMENTAL ***
//
// Returns any repository metadata that is
// useful within Taskcluster related services.
//
// See https://docs.taskcluster.net/reference/core/github/api-docs#repository
func (github *Github) Repository(owner, repo string) (*RepositoryResponse, error) {
	cd := tcclient.Client(*github)
	responseObject, _, err := (&cd).APICall(nil, "GET", "/repository/"+url.QueryEscape(owner)+"/"+url.QueryEscape(repo), new(RepositoryResponse), nil)
	return responseObject.(*RepositoryResponse), err
}

// Stability: *** EXPERIMENTAL ***
//
// For a given branch of a repository, this will always point
// to a status page for the most recent task triggered by that
// branch.
//
// Note: This is a redirect rather than a direct link.
//
// See https://docs.taskcluster.net/reference/core/github/api-docs#latest
func (github *Github) Latest(owner, repo, branch string) error {
	cd := tcclient.Client(*github)
	_, _, err := (&cd).APICall(nil, "GET", "/repository/"+url.QueryEscape(owner)+"/"+url.QueryEscape(repo)+"/"+url.QueryEscape(branch)+"/latest", nil, nil)
	return err
}

// Stability: *** EXPERIMENTAL ***
//
// For a given changeset (SHA) of a repository, this will attach a "commit status"
// on github. These statuses are links displayed next to each revision.
// The status is either OK (green check) or FAILURE (red cross),
// made of a custom title and link.
//
// Required scopes:
//   github:create-status:<owner>/<repo>
//
// See https://docs.taskcluster.net/reference/core/github/api-docs#createStatus
func (github *Github) CreateStatus(owner, repo, sha string, payload *CreateStatusRequest) error {
	cd := tcclient.Client(*github)
	_, _, err := (&cd).APICall(payload, "POST", "/repository/"+url.QueryEscape(owner)+"/"+url.QueryEscape(repo)+"/statuses/"+url.QueryEscape(sha), nil, nil)
	return err
}

// Stability: *** EXPERIMENTAL ***
//
// For a given Issue or Pull Request of a repository, this will write a new message.
//
// Required scopes:
//   github:create-comment:<owner>/<repo>
//
// See https://docs.taskcluster.net/reference/core/github/api-docs#createComment
func (github *Github) CreateComment(owner, repo, number string, payload *CreateCommentRequest) error {
	cd := tcclient.Client(*github)
	_, _, err := (&cd).APICall(payload, "POST", "/repository/"+url.QueryEscape(owner)+"/"+url.QueryEscape(repo)+"/issues/"+url.QueryEscape(number)+"/comments", nil, nil)
	return err
}
