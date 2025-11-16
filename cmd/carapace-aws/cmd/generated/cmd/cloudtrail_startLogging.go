package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_startLoggingCmd = &cobra.Command{
	Use:   "start-logging",
	Short: "Starts the recording of Amazon Web Services API calls and log file delivery for a trail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_startLoggingCmd).Standalone()

	cloudtrail_startLoggingCmd.Flags().String("name", "", "Specifies the name or the CloudTrail ARN of the trail for which CloudTrail logs Amazon Web Services API calls.")
	cloudtrail_startLoggingCmd.MarkFlagRequired("name")
	cloudtrailCmd.AddCommand(cloudtrail_startLoggingCmd)
}
