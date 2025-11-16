package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deleteResourceTreeCmd = &cobra.Command{
	Use:   "delete-resource-tree",
	Short: "Deletes an entire resource tree.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deleteResourceTreeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_deleteResourceTreeCmd).Standalone()

		forecast_deleteResourceTreeCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the parent resource to delete.")
		forecast_deleteResourceTreeCmd.MarkFlagRequired("resource-arn")
	})
	forecastCmd.AddCommand(forecast_deleteResourceTreeCmd)
}
