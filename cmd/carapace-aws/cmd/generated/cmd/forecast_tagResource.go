package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the specified tags to a resource with the specified `resourceArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_tagResourceCmd).Standalone()

		forecast_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource for which to list the tags.")
		forecast_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
		forecast_tagResourceCmd.MarkFlagRequired("resource-arn")
		forecast_tagResourceCmd.MarkFlagRequired("tags")
	})
	forecastCmd.AddCommand(forecast_tagResourceCmd)
}
