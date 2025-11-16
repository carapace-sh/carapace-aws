package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_deleteSnapshotCmd = &cobra.Command{
	Use:   "delete-snapshot",
	Short: "Deletes a snapshot from Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_deleteSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_deleteSnapshotCmd).Standalone()

		redshiftServerless_deleteSnapshotCmd.Flags().String("snapshot-name", "", "The name of the snapshot to be deleted.")
		redshiftServerless_deleteSnapshotCmd.MarkFlagRequired("snapshot-name")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_deleteSnapshotCmd)
}
