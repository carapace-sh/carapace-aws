package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_getCallAnalyticsJobCmd = &cobra.Command{
	Use:   "get-call-analytics-job",
	Short: "Provides information about the specified Call Analytics job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_getCallAnalyticsJobCmd).Standalone()

	transcribe_getCallAnalyticsJobCmd.Flags().String("call-analytics-job-name", "", "The name of the Call Analytics job you want information about.")
	transcribe_getCallAnalyticsJobCmd.MarkFlagRequired("call-analytics-job-name")
	transcribeCmd.AddCommand(transcribe_getCallAnalyticsJobCmd)
}
