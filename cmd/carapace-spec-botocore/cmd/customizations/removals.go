package customizations

import "slices"

var removals = map[string][]string{
	"ses": {
		"delete-verified-email-address",
		"list-verified-email-addresses",
		"verify-email-address",
	},
	"ec2": {
		"import-instance",
		"import-volume",
	},
	"emr": {
		"run-job-flow",
		"describe-job-flows",
		"add-job-flow-steps",
		"terminate-job-flows",
		"list-bootstrap-actions",
		"list-instance-groups",
		"set-termination-protection",
		"set-keep-job-flow-alive-when-no-steps",
		"set-visible-to-all-users",
		"set-unhealthy-node-replacement",
	},
	"kinesis": {
		"subscribe-to-shard",
	},
	"lexv2-runtime": {
		"start-conversation",
	},
	"lambda": {
		"invoke-with-response-stream",
	},
	"sagemaker-runtime": {
		"invoke-endpoint-with-response-stream",
	},
	"bedrock-runtime": {
		"invoke-model-with-bidirectional-stream",
		"invoke-model-with-response-stream",
		"converse-stream",
	},
	"bedrock-agent-runtime": {
		"invoke-agent",
		"invoke-flow",
		"invoke-inline-agent",
		"optimize-prompt",
		"retrieve-and-generate-stream",
	},
	"bedrock-agentcore": {
		"invoke-code-interpreter",
	},
	"qbusiness": {
		"chat",
	},
	"iotsitewise": {
		"invoke-assistant",
	},
	"logs": {
		"get-log-object",
	},
}

func Removal(service, operation string) bool {
	if operations, ok := removals[service]; ok {
		return slices.Contains(operations, operation)
	}
	return false
}
