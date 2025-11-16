package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_updateCallAnalyticsCategoryCmd = &cobra.Command{
	Use:   "update-call-analytics-category",
	Short: "Updates the specified Call Analytics category with new rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_updateCallAnalyticsCategoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_updateCallAnalyticsCategoryCmd).Standalone()

		transcribe_updateCallAnalyticsCategoryCmd.Flags().String("category-name", "", "The name of the Call Analytics category you want to update.")
		transcribe_updateCallAnalyticsCategoryCmd.Flags().String("input-type", "", "Choose whether you want to update a real-time or a post-call category.")
		transcribe_updateCallAnalyticsCategoryCmd.Flags().String("rules", "", "The rules used for the updated Call Analytics category.")
		transcribe_updateCallAnalyticsCategoryCmd.MarkFlagRequired("category-name")
		transcribe_updateCallAnalyticsCategoryCmd.MarkFlagRequired("rules")
	})
	transcribeCmd.AddCommand(transcribe_updateCallAnalyticsCategoryCmd)
}
