package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_describeReplicationConfigurationTemplatesCmd = &cobra.Command{
	Use:   "describe-replication-configuration-templates",
	Short: "Lists all ReplicationConfigurationTemplates, filtered by Source Server IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_describeReplicationConfigurationTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_describeReplicationConfigurationTemplatesCmd).Standalone()

		mgn_describeReplicationConfigurationTemplatesCmd.Flags().String("max-results", "", "Request to describe Replication Configuration template by max results.")
		mgn_describeReplicationConfigurationTemplatesCmd.Flags().String("next-token", "", "Request to describe Replication Configuration template by next token.")
		mgn_describeReplicationConfigurationTemplatesCmd.Flags().String("replication-configuration-template-ids", "", "Request to describe Replication Configuration template by template IDs.")
	})
	mgnCmd.AddCommand(mgn_describeReplicationConfigurationTemplatesCmd)
}
