package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_stopLoggingCmd = &cobra.Command{
	Use:   "stop-logging",
	Short: "Suspends the recording of Amazon Web Services API calls and log file delivery for the specified trail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_stopLoggingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_stopLoggingCmd).Standalone()

		cloudtrail_stopLoggingCmd.Flags().String("name", "", "Specifies the name or the CloudTrail ARN of the trail for which CloudTrail will stop logging Amazon Web Services API calls.")
		cloudtrail_stopLoggingCmd.MarkFlagRequired("name")
	})
	cloudtrailCmd.AddCommand(cloudtrail_stopLoggingCmd)
}
