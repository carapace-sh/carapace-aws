package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_modifyReplicationConfigCmd = &cobra.Command{
	Use:   "modify-replication-config",
	Short: "Modifies an existing DMS Serverless replication configuration that you can use to start a replication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_modifyReplicationConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_modifyReplicationConfigCmd).Standalone()

		dms_modifyReplicationConfigCmd.Flags().String("compute-config", "", "Configuration parameters for provisioning an DMS Serverless replication.")
		dms_modifyReplicationConfigCmd.Flags().String("replication-config-arn", "", "The Amazon Resource Name of the replication to modify.")
		dms_modifyReplicationConfigCmd.Flags().String("replication-config-identifier", "", "The new replication config to apply to the replication.")
		dms_modifyReplicationConfigCmd.Flags().String("replication-settings", "", "The settings for the replication.")
		dms_modifyReplicationConfigCmd.Flags().String("replication-type", "", "The type of replication.")
		dms_modifyReplicationConfigCmd.Flags().String("source-endpoint-arn", "", "The Amazon Resource Name (ARN) of the source endpoint for this DMS serverless replication configuration.")
		dms_modifyReplicationConfigCmd.Flags().String("supplemental-settings", "", "Additional settings for the replication.")
		dms_modifyReplicationConfigCmd.Flags().String("table-mappings", "", "Table mappings specified in the replication.")
		dms_modifyReplicationConfigCmd.Flags().String("target-endpoint-arn", "", "The Amazon Resource Name (ARN) of the target endpoint for this DMS serverless replication configuration.")
		dms_modifyReplicationConfigCmd.MarkFlagRequired("replication-config-arn")
	})
	dmsCmd.AddCommand(dms_modifyReplicationConfigCmd)
}
