package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_restoreTableCmd = &cobra.Command{
	Use:   "restore-table",
	Short: "Restores the table to the specified point in time within the `earliest_restorable_timestamp` and the current time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_restoreTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(keyspaces_restoreTableCmd).Standalone()

		keyspaces_restoreTableCmd.Flags().String("auto-scaling-specification", "", "The optional auto scaling settings for the restored table in provisioned capacity mode.")
		keyspaces_restoreTableCmd.Flags().String("capacity-specification-override", "", "Specifies the read/write throughput capacity mode for the target table.")
		keyspaces_restoreTableCmd.Flags().String("encryption-specification-override", "", "Specifies the encryption settings for the target table.")
		keyspaces_restoreTableCmd.Flags().String("point-in-time-recovery-override", "", "Specifies the `pointInTimeRecovery` settings for the target table.")
		keyspaces_restoreTableCmd.Flags().String("replica-specifications", "", "The optional Region specific settings of a multi-Regional table.")
		keyspaces_restoreTableCmd.Flags().String("restore-timestamp", "", "The restore timestamp in ISO 8601 format.")
		keyspaces_restoreTableCmd.Flags().String("source-keyspace-name", "", "The keyspace name of the source table.")
		keyspaces_restoreTableCmd.Flags().String("source-table-name", "", "The name of the source table.")
		keyspaces_restoreTableCmd.Flags().String("tags-override", "", "A list of key-value pair tags to be attached to the restored table.")
		keyspaces_restoreTableCmd.Flags().String("target-keyspace-name", "", "The name of the target keyspace.")
		keyspaces_restoreTableCmd.Flags().String("target-table-name", "", "The name of the target table.")
		keyspaces_restoreTableCmd.MarkFlagRequired("source-keyspace-name")
		keyspaces_restoreTableCmd.MarkFlagRequired("source-table-name")
		keyspaces_restoreTableCmd.MarkFlagRequired("target-keyspace-name")
		keyspaces_restoreTableCmd.MarkFlagRequired("target-table-name")
	})
	keyspacesCmd.AddCommand(keyspaces_restoreTableCmd)
}
