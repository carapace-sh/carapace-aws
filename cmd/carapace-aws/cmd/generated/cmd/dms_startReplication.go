package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startReplicationCmd = &cobra.Command{
	Use:   "start-replication",
	Short: "For a given DMS Serverless replication configuration, DMS connects to the source endpoint and collects the metadata to analyze the replication workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startReplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_startReplicationCmd).Standalone()

		dms_startReplicationCmd.Flags().String("cdc-start-position", "", "Indicates when you want a change data capture (CDC) operation to start.")
		dms_startReplicationCmd.Flags().String("cdc-start-time", "", "Indicates the start time for a change data capture (CDC) operation.")
		dms_startReplicationCmd.Flags().String("cdc-stop-position", "", "Indicates when you want a change data capture (CDC) operation to stop.")
		dms_startReplicationCmd.Flags().String("premigration-assessment-settings", "", "User-defined settings for the premigration assessment.")
		dms_startReplicationCmd.Flags().String("replication-config-arn", "", "The Amazon Resource Name of the replication for which to start replication.")
		dms_startReplicationCmd.Flags().String("start-replication-type", "", "The replication type.")
		dms_startReplicationCmd.MarkFlagRequired("replication-config-arn")
		dms_startReplicationCmd.MarkFlagRequired("start-replication-type")
	})
	dmsCmd.AddCommand(dms_startReplicationCmd)
}
