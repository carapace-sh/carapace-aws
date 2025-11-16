package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_stopApplicationCmd = &cobra.Command{
	Use:   "stop-application",
	Short: "Stops the application from processing data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_stopApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_stopApplicationCmd).Standalone()

		kinesisanalyticsv2_stopApplicationCmd.Flags().String("application-name", "", "The name of the running application to stop.")
		kinesisanalyticsv2_stopApplicationCmd.Flags().String("force", "", "Set to `true` to force the application to stop.")
		kinesisanalyticsv2_stopApplicationCmd.MarkFlagRequired("application-name")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_stopApplicationCmd)
}
