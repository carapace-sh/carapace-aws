package customizations

import (
	_ "embed"
	"slices"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed rds/aws.rds.add-option-to-option-group.yaml
var rds_addOptionToOptionGroup []byte

//go:embed rds/aws.rds.generate-db-auth-token.yaml
var rds_generateDbAuthToken []byte

//go:embed rds/aws.rds.remove-option-from-option-group.yaml
var rds_removeOptionFromOptionGroup []byte

func init() {
	customizations["rds"] = func(cmd *command.Command) error {
		cmd.Commands = slices.DeleteFunc(cmd.Commands, func(e command.Command) bool {
			return e.Name == "modify-option-group"
		})

		for _, spec := range [][]byte{
			rds_addOptionToOptionGroup,
			rds_generateDbAuthToken,
			rds_removeOptionFromOptionGroup,
		} {
			var specCommand command.Command
			if err := yaml.Unmarshal(spec, &specCommand); err != nil {
				return err
			}
			cmd.Commands = append(cmd.Commands, specCommand)
		}
		return nil
	}

	customizations["rds.create-db-instance-read-replica"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "source-region",
			Description: "The ID of the region that contains the source for the read replica",
			Value:       true,
		})
		return nil
	}
	customizations["rds.create-db-cluster"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "source-region",
			Description: "The ID of the region that contains the source for the db cluster",
			Value:       true,
		})
		return nil
	}
	customizations["rds.copy-db-snapshot"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "source-region",
			Description: "The ID of the region that contains the snapshot to be copied",
			Value:       true,
		})
		return nil
	}
	customizations["rds.copy-db-cluster-snapshot"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "source-region",
			Description: "The ID of the region that contains the snapshot to be copied",
			Value:       true,
		})
		return nil
	}
	customizations["rds.start-db-instance-automated-backups-replication"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "source-region",
			Description: "The ID of the region that contains the source for the db instance",
			Value:       true,
		})
		return nil
	}
}
