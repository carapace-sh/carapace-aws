package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

func init() {
	for _, operation := range []string{
		"mgn.associate-source-servers",
		"mgn.disassociate-source-servers",
		"mgn.start-cutover",
		"mgn.start-test",
		"mgn.terminate-target-instances",
	} {
		customizations[operation] = func(cmd *command.Command) error {
			cmd.AddFlag(command.Flag{
				Longhand:    "source-server-ids",
				Description: "Source server IDs list",
				Value:       true,
				Required:    true,
			})
			return nil
		}
	}

	for _, operation := range []string{
		"mgn.create-replication-configuration-template",
		"mgn.update-replication-configuration",
		"mgn.update-replication-configuration-template",
	} {
		customizations[operation] = func(cmd *command.Command) error {
			cmd.AddFlag(command.Flag{
				Longhand:    "replication-servers-security-groups-i-ds",
				Description: "Request to configure the Replication Server Security group ID during Replication Settings template creation",
				Value:       true,
				Required:    true,
				Nargs:       -1,
			})
			return nil
		}
	}

	customizations["mgn.describe-replication-configuration-templates"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "replication-configuration-template-i-ds",
			Description: "Request to describe Replication Configuration template by template IDs",
			Value:       true,
			Nargs:       -1,
		})
		return nil
	}
}
