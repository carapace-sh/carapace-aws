package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_updateSnapshotCmd = &cobra.Command{
	Use:   "update-snapshot",
	Short: "Updates a snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_updateSnapshotCmd).Standalone()

	redshiftServerless_updateSnapshotCmd.Flags().String("retention-period", "", "The new retention period of the snapshot.")
	redshiftServerless_updateSnapshotCmd.Flags().String("snapshot-name", "", "The name of the snapshot.")
	redshiftServerless_updateSnapshotCmd.MarkFlagRequired("snapshot-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_updateSnapshotCmd)
}
