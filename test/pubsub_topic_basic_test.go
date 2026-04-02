package test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

// TestPubSubTopicBasic tests creating a basic Pub/Sub topic.
func TestPubSubTopicBasic(t *testing.T) {
	t.Parallel()

	retrySleep := 5 * time.Second
	unique := strings.ToLower(random.UniqueId())
	projectID := mustEnv(t, "GOOGLE_CLOUD_PROJECT")
	region := "us-central1"

	tfOptions := &terraform.Options{
		TerraformDir: "..",
		NoColor:      true,
		Vars: map[string]interface{}{
			"environment":  "devl",
			"project_code": "tt",
			"region":       region,
			"pubsub_config": map[string]interface{}{
				"base_name": fmt.Sprintf("topic-%s", unique),
			},
		},
	}

	defer terraform.Destroy(t, tfOptions)
	terraform.InitAndApply(t, tfOptions)

	time.Sleep(retrySleep)

	expectedTopicName := fmt.Sprintf("tt-topic-%s-%s-devl", unique, region)
	outputTopicName := terraform.Output(t, tfOptions, "topic_name")
	require.Equal(t, expectedTopicName, outputTopicName)

	client := newPubSubClient(t, projectID)
	defer client.Close()
	require.True(t, topicExists(t, client, outputTopicName))
}
