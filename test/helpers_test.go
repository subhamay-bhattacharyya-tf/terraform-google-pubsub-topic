// File: test/helpers_test.go
package test

import (
	"context"
	"os"
	"strings"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/require"
)

// mustEnv retrieves a required environment variable, failing the test if absent.
func mustEnv(t *testing.T, key string) string {
	t.Helper()
	v := strings.TrimSpace(os.Getenv(key))
	require.NotEmpty(t, v, "Missing required environment variable %s", key)
	return v
}

// newPubSubClient creates an authenticated Pub/Sub client.
func newPubSubClient(t *testing.T, projectID string) *pubsub.Client {
	t.Helper()
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	require.NoError(t, err, "Failed to create Pub/Sub client")
	return client
}

// topicExists reports whether the named topic is accessible.
func topicExists(t *testing.T, client *pubsub.Client, topicID string) bool {
	t.Helper()
	ctx := context.Background()
	topic := client.Topic(topicID)
	exists, err := topic.Exists(ctx)
	require.NoError(t, err, "Failed to check topic existence for %s", topicID)
	return exists
}
