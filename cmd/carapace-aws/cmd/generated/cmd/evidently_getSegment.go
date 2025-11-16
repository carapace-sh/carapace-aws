package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_getSegmentCmd = &cobra.Command{
	Use:   "get-segment",
	Short: "Returns information about the specified segment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_getSegmentCmd).Standalone()

	evidently_getSegmentCmd.Flags().String("segment", "", "The ARN of the segment to return information for.")
	evidently_getSegmentCmd.MarkFlagRequired("segment")
	evidentlyCmd.AddCommand(evidently_getSegmentCmd)
}
