package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectContactLens_listRealtimeContactAnalysisSegmentsCmd = &cobra.Command{
	Use:   "list-realtime-contact-analysis-segments",
	Short: "Provides a list of analysis segments for a real-time analysis session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectContactLens_listRealtimeContactAnalysisSegmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectContactLens_listRealtimeContactAnalysisSegmentsCmd).Standalone()

		connectContactLens_listRealtimeContactAnalysisSegmentsCmd.Flags().String("contact-id", "", "The identifier of the contact.")
		connectContactLens_listRealtimeContactAnalysisSegmentsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connectContactLens_listRealtimeContactAnalysisSegmentsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connectContactLens_listRealtimeContactAnalysisSegmentsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connectContactLens_listRealtimeContactAnalysisSegmentsCmd.MarkFlagRequired("contact-id")
		connectContactLens_listRealtimeContactAnalysisSegmentsCmd.MarkFlagRequired("instance-id")
	})
	connectContactLensCmd.AddCommand(connectContactLens_listRealtimeContactAnalysisSegmentsCmd)
}
