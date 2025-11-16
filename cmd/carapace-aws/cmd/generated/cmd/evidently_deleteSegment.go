package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_deleteSegmentCmd = &cobra.Command{
	Use:   "delete-segment",
	Short: "Deletes a segment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_deleteSegmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_deleteSegmentCmd).Standalone()

		evidently_deleteSegmentCmd.Flags().String("segment", "", "Specifies the segment to delete.")
		evidently_deleteSegmentCmd.MarkFlagRequired("segment")
	})
	evidentlyCmd.AddCommand(evidently_deleteSegmentCmd)
}
