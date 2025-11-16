package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getInsightImpactGraphCmd = &cobra.Command{
	Use:   "get-insight-impact-graph",
	Short: "Retrieves a service graph structure filtered by the specified insight.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getInsightImpactGraphCmd).Standalone()

	xray_getInsightImpactGraphCmd.Flags().String("end-time", "", "The estimated end time of the insight, in Unix time seconds.")
	xray_getInsightImpactGraphCmd.Flags().String("insight-id", "", "The insight's unique identifier.")
	xray_getInsightImpactGraphCmd.Flags().String("next-token", "", "Specify the pagination token returned by a previous request to retrieve the next page of results.")
	xray_getInsightImpactGraphCmd.Flags().String("start-time", "", "The estimated start time of the insight, in Unix time seconds.")
	xray_getInsightImpactGraphCmd.MarkFlagRequired("end-time")
	xray_getInsightImpactGraphCmd.MarkFlagRequired("insight-id")
	xray_getInsightImpactGraphCmd.MarkFlagRequired("start-time")
	xrayCmd.AddCommand(xray_getInsightImpactGraphCmd)
}
