package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_describeReplicationConfigurationTemplatesCmd = &cobra.Command{
	Use:   "describe-replication-configuration-templates",
	Short: "Lists all ReplicationConfigurationTemplates, filtered by Source Server IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_describeReplicationConfigurationTemplatesCmd).Standalone()

	drs_describeReplicationConfigurationTemplatesCmd.Flags().String("max-results", "", "Maximum number of Replication Configuration Templates to retrieve.")
	drs_describeReplicationConfigurationTemplatesCmd.Flags().String("next-token", "", "The token of the next Replication Configuration Template to retrieve.")
	drs_describeReplicationConfigurationTemplatesCmd.Flags().String("replication-configuration-template-ids", "", "The IDs of the Replication Configuration Templates to retrieve.")
	drsCmd.AddCommand(drs_describeReplicationConfigurationTemplatesCmd)
}
