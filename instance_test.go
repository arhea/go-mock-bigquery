package mockbigquery_test

import (
	"context"
	"testing"

	mockbigquery "github.com/arhea/go-mock-bigquery"
	"github.com/stretchr/testify/assert"
)

func TestInstance(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	const projectID = "test-project"
	const datasetID = "test-dataset"

	mock, err := mockbigquery.NewInstance(ctx, t, projectID, datasetID)

	if err != nil {
		t.Fatalf("creating the instance: %v", err)
		return
	}

	// close the mock
	defer mock.Close(ctx)

	port, err := mock.Port(ctx)

	if err != nil {
		t.Fatalf("getting the mapped port of the instance: %v", err)
		return
	}

	if port.Port() == "" {
		t.Fatalf("port should not be empty")
		return
	}

	assert.Equal(t, mock.ProjectID(), projectID)
	assert.Equal(t, mock.DatasetID(), datasetID)
}
