package mockbigquery_test

import (
	"context"
	"fmt"
	"testing"

	mockbigquery "github.com/arhea/go-mock-bigquery"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	const projectID = "test-project"
	const datasetID = "test-dataset"

	mock, err := mockbigquery.NewClient(ctx, t, projectID, datasetID)

	if err != nil {
		t.Fatalf("creating the client: %v", err)
		return
	}

	// close the mock
	defer mock.Close(ctx)

	assert.NotNil(t, mock.Client(), "client should not be nil")
	assert.Equal(t, projectID, mock.ProjectID())
	assert.Equal(t, datasetID, mock.DatasetID())
	assert.Equal(t, fmt.Sprintf("%s.%s", projectID, datasetID), mock.FullName())
}
