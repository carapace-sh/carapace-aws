package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_deleteReplicationConfigurationCmd = &cobra.Command{
	Use:   "delete-replication-configuration",
	Short: "Deletes a replication configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_deleteReplicationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_deleteReplicationConfigurationCmd).Standalone()

		efs_deleteReplicationConfigurationCmd.Flags().String("deletion-mode", "", "When replicating across Amazon Web Services accounts or across Amazon Web Services Regions, Amazon EFS deletes the replication configuration from both the source and destination account or Region (`ALL_CONFIGURATIONS`) by default.")
		efs_deleteReplicationConfigurationCmd.Flags().String("source-file-system-id", "", "The ID of the source file system in the replication configuration.")
		efs_deleteReplicationConfigurationCmd.MarkFlagRequired("source-file-system-id")
	})
	efsCmd.AddCommand(efs_deleteReplicationConfigurationCmd)
}
