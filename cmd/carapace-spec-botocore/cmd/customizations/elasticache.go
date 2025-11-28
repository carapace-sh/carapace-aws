package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

func init() {
	customizations["elasticache.create-replication-group"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "preferred-cache-cluster-azs",
			Description: "A list of EC2 Availability Zones in which the replication group’s clusters are created",
			Value:       true,
			Nargs:       -1,
		})
		return nil
	}
}
