package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteReplicationConfigCmd = &cobra.Command{
	Use:   "delete-replication-config",
	Short: "Deletes an DMS Serverless replication configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteReplicationConfigCmd).Standalone()

	dms_deleteReplicationConfigCmd.Flags().String("replication-config-arn", "", "The replication config to delete.")
	dms_deleteReplicationConfigCmd.MarkFlagRequired("replication-config-arn")
	dmsCmd.AddCommand(dms_deleteReplicationConfigCmd)
}
