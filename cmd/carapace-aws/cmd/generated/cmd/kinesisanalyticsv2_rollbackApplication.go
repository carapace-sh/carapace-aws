package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_rollbackApplicationCmd = &cobra.Command{
	Use:   "rollback-application",
	Short: "Reverts the application to the previous running version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_rollbackApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_rollbackApplicationCmd).Standalone()

		kinesisanalyticsv2_rollbackApplicationCmd.Flags().String("application-name", "", "The name of the application.")
		kinesisanalyticsv2_rollbackApplicationCmd.Flags().String("current-application-version-id", "", "The current application version ID.")
		kinesisanalyticsv2_rollbackApplicationCmd.MarkFlagRequired("application-name")
		kinesisanalyticsv2_rollbackApplicationCmd.MarkFlagRequired("current-application-version-id")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_rollbackApplicationCmd)
}
