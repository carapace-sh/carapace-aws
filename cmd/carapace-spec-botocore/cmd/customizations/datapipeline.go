package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed datapipeline/aws.datapipeline.create-default-roles.yaml
var datapipeline_createDefaultRoles []byte

//go:embed datapipeline/aws.datapipeline.list-runs.yaml
var datapipeline_listRuns []byte

func init() {
	customizations["datapipeline"] = func(cmd *command.Command) error {
		for _, spec := range [][]byte{
			datapipeline_createDefaultRoles,
			datapipeline_listRuns,
		} {
			var specCommand command.Command
			if err := yaml.Unmarshal(spec, &specCommand); err != nil {
				return err
			}
			cmd.Commands = append(cmd.Commands, specCommand)
		}
		return nil
	}

	customizations["datapipeline.activate-pipeline"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "parameter-values-uri",
			Description: "The JSON parameter values.",
			Value:       true,
		})
		return nil
	}

	customizations["datapipeline.put-pipeline-definition"] = func(cmd *command.Command) error {
		delete(cmd.Flags, "--pipeline-objects=!")

		cmd.AddFlag(command.Flag{
			Longhand:    "parameter-values-uri",
			Description: "The JSON parameter values.",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "pipeline-definition",
			Description: "The JSON pipeline definition.",
			Value:       true,
			Required:    true,
		})
		return nil
	}
}
