package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_listApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalytics_listApplicationsCmd).Standalone()

		kinesisanalytics_listApplicationsCmd.Flags().String("exclusive-start-application-name", "", "Name of the application to start the list with.")
		kinesisanalytics_listApplicationsCmd.Flags().String("limit", "", "Maximum number of applications to list.")
	})
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_listApplicationsCmd)
}
