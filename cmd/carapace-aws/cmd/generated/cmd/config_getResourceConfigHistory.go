package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getResourceConfigHistoryCmd = &cobra.Command{
	Use:   "get-resource-config-history",
	Short: "For accurate reporting on the compliance status, you must record the `AWS::Config::ResourceCompliance` resource type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getResourceConfigHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_getResourceConfigHistoryCmd).Standalone()

		config_getResourceConfigHistoryCmd.Flags().String("chronological-order", "", "The chronological order for configuration items listed.")
		config_getResourceConfigHistoryCmd.Flags().String("earlier-time", "", "The chronologically earliest time in the time range for which the history requested.")
		config_getResourceConfigHistoryCmd.Flags().String("later-time", "", "The chronologically latest time in the time range for which the history requested.")
		config_getResourceConfigHistoryCmd.Flags().String("limit", "", "The maximum number of configuration items returned on each page.")
		config_getResourceConfigHistoryCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_getResourceConfigHistoryCmd.Flags().String("resource-id", "", "The ID of the resource (for example., `sg-xxxxxx`).")
		config_getResourceConfigHistoryCmd.Flags().String("resource-type", "", "The resource type.")
		config_getResourceConfigHistoryCmd.MarkFlagRequired("resource-id")
		config_getResourceConfigHistoryCmd.MarkFlagRequired("resource-type")
	})
	configCmd.AddCommand(config_getResourceConfigHistoryCmd)
}
