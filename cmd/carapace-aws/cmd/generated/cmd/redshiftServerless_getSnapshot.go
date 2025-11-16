package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getSnapshotCmd = &cobra.Command{
	Use:   "get-snapshot",
	Short: "Returns information about a specific snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getSnapshotCmd).Standalone()

	redshiftServerless_getSnapshotCmd.Flags().String("owner-account", "", "The owner Amazon Web Services account of a snapshot shared with another user.")
	redshiftServerless_getSnapshotCmd.Flags().String("snapshot-arn", "", "The Amazon Resource Name (ARN) of the snapshot to return.")
	redshiftServerless_getSnapshotCmd.Flags().String("snapshot-name", "", "The name of the snapshot to return.")
	redshiftServerlessCmd.AddCommand(redshiftServerless_getSnapshotCmd)
}
