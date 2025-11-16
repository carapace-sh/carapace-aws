package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listRealtimeContactAnalysisSegmentsV2Cmd = &cobra.Command{
	Use:   "list-realtime-contact-analysis-segments-v2",
	Short: "Provides a list of analysis segments for a real-time chat analysis session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listRealtimeContactAnalysisSegmentsV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listRealtimeContactAnalysisSegmentsV2Cmd).Standalone()

		connect_listRealtimeContactAnalysisSegmentsV2Cmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
		connect_listRealtimeContactAnalysisSegmentsV2Cmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listRealtimeContactAnalysisSegmentsV2Cmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listRealtimeContactAnalysisSegmentsV2Cmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listRealtimeContactAnalysisSegmentsV2Cmd.Flags().String("output-type", "", "The Contact Lens output type to be returned.")
		connect_listRealtimeContactAnalysisSegmentsV2Cmd.Flags().String("segment-types", "", "Enum with segment types .")
		connect_listRealtimeContactAnalysisSegmentsV2Cmd.MarkFlagRequired("contact-id")
		connect_listRealtimeContactAnalysisSegmentsV2Cmd.MarkFlagRequired("instance-id")
		connect_listRealtimeContactAnalysisSegmentsV2Cmd.MarkFlagRequired("output-type")
		connect_listRealtimeContactAnalysisSegmentsV2Cmd.MarkFlagRequired("segment-types")
	})
	connectCmd.AddCommand(connect_listRealtimeContactAnalysisSegmentsV2Cmd)
}
