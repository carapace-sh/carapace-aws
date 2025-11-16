package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_stopResourceCmd = &cobra.Command{
	Use:   "stop-resource",
	Short: "Stops a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_stopResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_stopResourceCmd).Standalone()

		forecast_stopResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource to stop.")
		forecast_stopResourceCmd.MarkFlagRequired("resource-arn")
	})
	forecastCmd.AddCommand(forecast_stopResourceCmd)
}
