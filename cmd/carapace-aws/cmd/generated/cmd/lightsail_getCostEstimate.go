package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getCostEstimateCmd = &cobra.Command{
	Use:   "get-cost-estimate",
	Short: "Retrieves information about the cost estimate for a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getCostEstimateCmd).Standalone()

	lightsail_getCostEstimateCmd.Flags().String("end-time", "", "The cost estimate end time.")
	lightsail_getCostEstimateCmd.Flags().String("resource-name", "", "The resource name.")
	lightsail_getCostEstimateCmd.Flags().String("start-time", "", "The cost estimate start time.")
	lightsail_getCostEstimateCmd.MarkFlagRequired("end-time")
	lightsail_getCostEstimateCmd.MarkFlagRequired("resource-name")
	lightsail_getCostEstimateCmd.MarkFlagRequired("start-time")
	lightsailCmd.AddCommand(lightsail_getCostEstimateCmd)
}
