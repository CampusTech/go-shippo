package shippo

import "github.com/campustech/go-shippo/client"

const (
	APIVersion20140211 = "2014-02-11"
	APIVesrion20161025 = "2016-10-25"
	APIVersion20170329 = "2017-03-29"
	APIVersion20180208 = "2018-02-08"
)

// NewClient creates a new Shippo client.
func NewClient(privateToken string) *client.Client {
	return client.NewClient(privateToken, "", "")
}

// NewClientWithVersion creates a new Shippo client with API version explicitly specified.
func NewClientWithVersion(privateToken, apiVersion string) *client.Client {
	return client.NewClient(privateToken, apiVersion, "")
}

// NewClientWithPlatformAccountID creates a new Shippo client with platform account ID explicitly specified.
func NewClientWithPlatformAccountID(privateToken, platformAccountID string) *client.Client {
	return client.NewClient(privateToken, "", platformAccountID)
}

// NewClientWithVersionAndPlatformAccountID creates a new Shippo client with API version and platform account ID explicitly specified.
func NewClientWithVersionAndPlatformAccountID(privateToken, apiVersion, platformAccountID string) *client.Client {
	return client.NewClient(privateToken, apiVersion, platformAccountID)
}
