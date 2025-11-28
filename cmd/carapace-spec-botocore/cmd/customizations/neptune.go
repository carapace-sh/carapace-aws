package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

func init() {
	customizations["neptune.create-db-cluster"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "source-region",
			Description: "The ID of the region that contains the source for the db cluster",
			Value:       true,
		})
		return nil
	}
	customizations["neptune.copy-db-cluster-snapshot"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "source-region",
			Description: "The ID of the region that contains the snapshot to be copied",
			Value:       true,
		})
		return nil
	}
}
