package mockbigquery

import (
	"context"
	"fmt"
	"testing"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

// Client wraps the Redis client and the underlying instance for testing. It provides helper methods for
// accessing the client, closing the client, and closing the instance.
type Client struct {
	t        *testing.T
	instance *Instance
	client   *bigquery.Client
}

// Client returns the underlying Redis client.
func (r *Client) Client() *bigquery.Client {
	r.t.Helper()

	return r.client
}

// FullName returns the fullly qualified dataset name.
func (c *Client) FullName() string {
	return fmt.Sprintf("%s.%s", c.ProjectID(), c.DatasetID())
}

// ProjectID returns the project id of the underlying instance.
func (c *Client) ProjectID() string {
	return c.instance.ProjectID()
}

// DatasetID returns the dataset id of the underlying instance.
func (c *Client) DatasetID() string {
	return c.instance.DatasetID()
}

// Close closes the underlying BigQuery client and the instance.
func (r *Client) Close(ctx context.Context) {
	r.t.Helper()

	if err := r.client.Close(); err != nil {
		r.t.Logf("closing big query client: %v", err)
	}

	r.instance.Close(ctx)
}

// NewClient creates a new BigQuery client that connects to an underlying emulator instance for testing.
func NewClient(ctx context.Context, t *testing.T, projectID string, datasetID string) (*Client, error) {
	t.Helper()

	instance, err := NewInstance(ctx, t, projectID, datasetID)

	if err != nil {
		return nil, fmt.Errorf("creating the instance: %v", err)
	}

	bigQueryPort, err := instance.Port(ctx)

	if err != nil {
		return nil, fmt.Errorf("getting the mapped port of the instance: %v", err)
	}

	time.Sleep(time.Millisecond * 500)

	// create the mock client
	bqClient, err := bigquery.NewClient(
		ctx,
		projectID,
		option.WithEndpoint("http://localhost:"+bigQueryPort.Port()),
		option.WithoutAuthentication(),
	)

	if err != nil {
		t.Errorf("error creating mock client: %v", err)
		return nil, err
	}

	client := &Client{
		t:        t,
		instance: instance,
		client:   bqClient,
	}

	return client, nil
}
