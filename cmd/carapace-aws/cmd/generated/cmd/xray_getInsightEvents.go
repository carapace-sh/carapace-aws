package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getInsightEventsCmd = &cobra.Command{
	Use:   "get-insight-events",
	Short: "X-Ray reevaluates insights periodically until they're resolved, and records each intermediate state as an event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getInsightEventsCmd).Standalone()

	xray_getInsightEventsCmd.Flags().String("insight-id", "", "The insight's unique identifier.")
	xray_getInsightEventsCmd.Flags().String("max-results", "", "Used to retrieve at most the specified value of events.")
	xray_getInsightEventsCmd.Flags().String("next-token", "", "Specify the pagination token returned by a previous request to retrieve the next page of events.")
	xray_getInsightEventsCmd.MarkFlagRequired("insight-id")
	xrayCmd.AddCommand(xray_getInsightEventsCmd)
}
