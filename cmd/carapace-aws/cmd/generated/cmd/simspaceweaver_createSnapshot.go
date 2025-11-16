package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_createSnapshotCmd = &cobra.Command{
	Use:   "create-snapshot",
	Short: "Creates a snapshot of the specified simulation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_createSnapshotCmd).Standalone()

	simspaceweaver_createSnapshotCmd.Flags().String("destination", "", "The Amazon S3 bucket and optional folder (object key prefix) where SimSpace Weaver creates the snapshot file.")
	simspaceweaver_createSnapshotCmd.Flags().String("simulation", "", "The name of the simulation.")
	simspaceweaver_createSnapshotCmd.MarkFlagRequired("destination")
	simspaceweaver_createSnapshotCmd.MarkFlagRequired("simulation")
	simspaceweaverCmd.AddCommand(simspaceweaver_createSnapshotCmd)
}
