package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_deleteApplicationOutputCmd = &cobra.Command{
	Use:   "delete-application-output",
	Short: "Deletes the output destination configuration from your SQL-based Kinesis Data Analytics application's configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_deleteApplicationOutputCmd).Standalone()

	kinesisanalyticsv2_deleteApplicationOutputCmd.Flags().String("application-name", "", "The application name.")
	kinesisanalyticsv2_deleteApplicationOutputCmd.Flags().String("current-application-version-id", "", "The application version.")
	kinesisanalyticsv2_deleteApplicationOutputCmd.Flags().String("output-id", "", "The ID of the configuration to delete.")
	kinesisanalyticsv2_deleteApplicationOutputCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2_deleteApplicationOutputCmd.MarkFlagRequired("current-application-version-id")
	kinesisanalyticsv2_deleteApplicationOutputCmd.MarkFlagRequired("output-id")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_deleteApplicationOutputCmd)
}
