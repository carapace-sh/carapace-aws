package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes the specified tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_untagResourceCmd).Standalone()

		forecast_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource for which to list the tags.")
		forecast_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to be removed.")
		forecast_untagResourceCmd.MarkFlagRequired("resource-arn")
		forecast_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	forecastCmd.AddCommand(forecast_untagResourceCmd)
}
