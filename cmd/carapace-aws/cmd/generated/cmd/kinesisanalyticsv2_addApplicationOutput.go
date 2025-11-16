package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_addApplicationOutputCmd = &cobra.Command{
	Use:   "add-application-output",
	Short: "Adds an external destination to your SQL-based Kinesis Data Analytics application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_addApplicationOutputCmd).Standalone()

	kinesisanalyticsv2_addApplicationOutputCmd.Flags().String("application-name", "", "The name of the application to which you want to add the output configuration.")
	kinesisanalyticsv2_addApplicationOutputCmd.Flags().String("current-application-version-id", "", "The version of the application to which you want to add the output configuration.")
	kinesisanalyticsv2_addApplicationOutputCmd.Flags().String("output", "", "An array of objects, each describing one output configuration.")
	kinesisanalyticsv2_addApplicationOutputCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2_addApplicationOutputCmd.MarkFlagRequired("current-application-version-id")
	kinesisanalyticsv2_addApplicationOutputCmd.MarkFlagRequired("output")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_addApplicationOutputCmd)
}
