package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_stopReplicationCmd = &cobra.Command{
	Use:   "stop-replication",
	Short: "For a given DMS Serverless replication configuration, DMS stops any and all ongoing DMS Serverless replications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_stopReplicationCmd).Standalone()

	dms_stopReplicationCmd.Flags().String("replication-config-arn", "", "The Amazon Resource Name of the replication to stop.")
	dms_stopReplicationCmd.MarkFlagRequired("replication-config-arn")
	dmsCmd.AddCommand(dms_stopReplicationCmd)
}
