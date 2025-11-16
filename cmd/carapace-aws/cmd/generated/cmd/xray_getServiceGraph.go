package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getServiceGraphCmd = &cobra.Command{
	Use:   "get-service-graph",
	Short: "Retrieves a document that describes services that process incoming requests, and downstream services that they call as a result.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getServiceGraphCmd).Standalone()

	xray_getServiceGraphCmd.Flags().String("end-time", "", "The end of the timeframe for which to generate a graph.")
	xray_getServiceGraphCmd.Flags().String("group-arn", "", "The Amazon Resource Name (ARN) of a group based on which you want to generate a graph.")
	xray_getServiceGraphCmd.Flags().String("group-name", "", "The name of a group based on which you want to generate a graph.")
	xray_getServiceGraphCmd.Flags().String("next-token", "", "Pagination token.")
	xray_getServiceGraphCmd.Flags().String("start-time", "", "The start of the time frame for which to generate a graph.")
	xray_getServiceGraphCmd.MarkFlagRequired("end-time")
	xray_getServiceGraphCmd.MarkFlagRequired("start-time")
	xrayCmd.AddCommand(xray_getServiceGraphCmd)
}
