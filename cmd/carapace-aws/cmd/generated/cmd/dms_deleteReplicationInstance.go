package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteReplicationInstanceCmd = &cobra.Command{
	Use:   "delete-replication-instance",
	Short: "Deletes the specified replication instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteReplicationInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_deleteReplicationInstanceCmd).Standalone()

		dms_deleteReplicationInstanceCmd.Flags().String("replication-instance-arn", "", "The Amazon Resource Name (ARN) of the replication instance to be deleted.")
		dms_deleteReplicationInstanceCmd.MarkFlagRequired("replication-instance-arn")
	})
	dmsCmd.AddCommand(dms_deleteReplicationInstanceCmd)
}
