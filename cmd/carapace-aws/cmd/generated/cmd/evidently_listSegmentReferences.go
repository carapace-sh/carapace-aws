package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_listSegmentReferencesCmd = &cobra.Command{
	Use:   "list-segment-references",
	Short: "Use this operation to find which experiments or launches are using a specified segment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_listSegmentReferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_listSegmentReferencesCmd).Standalone()

		evidently_listSegmentReferencesCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		evidently_listSegmentReferencesCmd.Flags().String("next-token", "", "The token to use when requesting the next set of results.")
		evidently_listSegmentReferencesCmd.Flags().String("segment", "", "The ARN of the segment that you want to view information for.")
		evidently_listSegmentReferencesCmd.Flags().String("type", "", "Specifies whether to return information about launches or experiments that use this segment.")
		evidently_listSegmentReferencesCmd.MarkFlagRequired("segment")
		evidently_listSegmentReferencesCmd.MarkFlagRequired("type")
	})
	evidentlyCmd.AddCommand(evidently_listSegmentReferencesCmd)
}
