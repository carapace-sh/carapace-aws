package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_deleteCallAnalyticsJobCmd = &cobra.Command{
	Use:   "delete-call-analytics-job",
	Short: "Deletes a Call Analytics job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_deleteCallAnalyticsJobCmd).Standalone()

	transcribe_deleteCallAnalyticsJobCmd.Flags().String("call-analytics-job-name", "", "The name of the Call Analytics job you want to delete.")
	transcribe_deleteCallAnalyticsJobCmd.MarkFlagRequired("call-analytics-job-name")
	transcribeCmd.AddCommand(transcribe_deleteCallAnalyticsJobCmd)
}
