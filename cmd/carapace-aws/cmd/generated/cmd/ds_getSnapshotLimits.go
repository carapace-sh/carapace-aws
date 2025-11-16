package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_getSnapshotLimitsCmd = &cobra.Command{
	Use:   "get-snapshot-limits",
	Short: "Obtains the manual snapshot limits for a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_getSnapshotLimitsCmd).Standalone()

	ds_getSnapshotLimitsCmd.Flags().String("directory-id", "", "Contains the identifier of the directory to obtain the limits for.")
	ds_getSnapshotLimitsCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_getSnapshotLimitsCmd)
}
