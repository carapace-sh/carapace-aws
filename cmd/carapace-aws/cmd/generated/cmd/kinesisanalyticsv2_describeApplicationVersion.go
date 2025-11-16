package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_describeApplicationVersionCmd = &cobra.Command{
	Use:   "describe-application-version",
	Short: "Provides a detailed description of a specified version of the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_describeApplicationVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_describeApplicationVersionCmd).Standalone()

		kinesisanalyticsv2_describeApplicationVersionCmd.Flags().String("application-name", "", "The name of the application for which you want to get the version description.")
		kinesisanalyticsv2_describeApplicationVersionCmd.Flags().String("application-version-id", "", "The ID of the application version for which you want to get the description.")
		kinesisanalyticsv2_describeApplicationVersionCmd.MarkFlagRequired("application-name")
		kinesisanalyticsv2_describeApplicationVersionCmd.MarkFlagRequired("application-version-id")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_describeApplicationVersionCmd)
}
