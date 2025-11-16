package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_addApplicationInputCmd = &cobra.Command{
	Use:   "add-application-input",
	Short: "Adds a streaming source to your SQL-based Kinesis Data Analytics application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_addApplicationInputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_addApplicationInputCmd).Standalone()

		kinesisanalyticsv2_addApplicationInputCmd.Flags().String("application-name", "", "The name of your existing application to which you want to add the streaming source.")
		kinesisanalyticsv2_addApplicationInputCmd.Flags().String("current-application-version-id", "", "The current version of your application.")
		kinesisanalyticsv2_addApplicationInputCmd.Flags().String("input", "", "The [Input]() to add.")
		kinesisanalyticsv2_addApplicationInputCmd.MarkFlagRequired("application-name")
		kinesisanalyticsv2_addApplicationInputCmd.MarkFlagRequired("current-application-version-id")
		kinesisanalyticsv2_addApplicationInputCmd.MarkFlagRequired("input")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_addApplicationInputCmd)
}
