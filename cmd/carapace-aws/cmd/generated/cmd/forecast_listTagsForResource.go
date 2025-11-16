package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for an Amazon Forecast resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_listTagsForResourceCmd).Standalone()

		forecast_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource for which to list the tags.")
		forecast_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	forecastCmd.AddCommand(forecast_listTagsForResourceCmd)
}
