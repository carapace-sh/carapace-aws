package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

func init() {
	customizations["bedrock-agentcore.invoke-agent-runtime-command"] = func(cmd *command.Command) error {
		return &SkipError{}
	}
	customizations["bedrock-agentcore.invoke-harness"] = func(cmd *command.Command) error {
		return &SkipError{}
	}
}
