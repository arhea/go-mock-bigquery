# Mock BigQuery Emulator

![Tests](https://github.com/arhea/go-mock-bigquery/actions/workflows/main.yml/badge.svg?branch=main) ![goreportcard](https://goreportcard.com/badge/github.com/arhea/go-mock-bigquery)

Provide a mock BigQuery Emulator instance and optionally a mock BigQuery client for testing purposes. This library is built so you can mock BigQuery using the [bigquery-emulator](https://github.com/goccy/bigquery-emulator) project. You will need to have Docker running on your local machine or within your CI environment.

This library is built on top of [testcontainers](https://testcontainers.com/).

## Usage

Creating a mock instance for creating a customer connection.

```golang
func TestXXX(t *testing.T) {
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

	// ... my test code
}
```

Creating a mock redis client for interacting with Redis.

```golang
func TestXXX(t *testing.T) {
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

    bqClient := mock.Client()

	// ... my test code
}
```
