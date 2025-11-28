package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed logs/aws.logs.tail.yaml
var logs_tail []byte

func init() {
	customizations["logs"] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(logs_tail, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}

	customizations["logs.start-live-tail"] = func(cmd *command.Command) error {
		// TODO these shouldn't have been added - a condition still missing? or is it an edge case in awscli?
		delete(cmd.Flags, "--generate-cli-skeleton")
		delete(cmd.Flags, "--cli-input-json=")
		delete(cmd.Flags, "--cli-input-yaml=")

		cmd.AddFlag(command.Flag{
			Longhand:    "mode",
			Description: "Live Tail mode",
			Value:       true,
		})
		if cmd.Completion.Flag == nil {
			cmd.Completion.Flag = map[string][]string{}
		}
		cmd.Completion.Flag["mode"] = []string{
			"interative\tStarts a Live Tail session in interactive mode",
			"print-only\tStarts a Live Tail session in print-only mode",
		}
		return nil
	}
}
