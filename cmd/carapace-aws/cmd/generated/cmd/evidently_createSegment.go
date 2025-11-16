package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_createSegmentCmd = &cobra.Command{
	Use:   "create-segment",
	Short: "Use this operation to define a *segment* of your audience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_createSegmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_createSegmentCmd).Standalone()

		evidently_createSegmentCmd.Flags().String("description", "", "An optional description for this segment.")
		evidently_createSegmentCmd.Flags().String("name", "", "A name for the segment.")
		evidently_createSegmentCmd.Flags().String("pattern", "", "The pattern to use for the segment.")
		evidently_createSegmentCmd.Flags().String("tags", "", "Assigns one or more tags (key-value pairs) to the segment.")
		evidently_createSegmentCmd.MarkFlagRequired("name")
		evidently_createSegmentCmd.MarkFlagRequired("pattern")
	})
	evidentlyCmd.AddCommand(evidently_createSegmentCmd)
}
