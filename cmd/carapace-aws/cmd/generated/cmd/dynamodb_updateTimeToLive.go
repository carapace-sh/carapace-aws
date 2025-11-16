package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_updateTimeToLiveCmd = &cobra.Command{
	Use:   "update-time-to-live",
	Short: "The `UpdateTimeToLive` method enables or disables Time to Live (TTL) for the specified table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_updateTimeToLiveCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_updateTimeToLiveCmd).Standalone()

		dynamodb_updateTimeToLiveCmd.Flags().String("table-name", "", "The name of the table to be configured.")
		dynamodb_updateTimeToLiveCmd.Flags().String("time-to-live-specification", "", "Represents the settings used to enable or disable Time to Live for the specified table.")
		dynamodb_updateTimeToLiveCmd.MarkFlagRequired("table-name")
		dynamodb_updateTimeToLiveCmd.MarkFlagRequired("time-to-live-specification")
	})
	dynamodbCmd.AddCommand(dynamodb_updateTimeToLiveCmd)
}
