package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_getCallAnalyticsCategoryCmd = &cobra.Command{
	Use:   "get-call-analytics-category",
	Short: "Provides information about the specified Call Analytics category.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_getCallAnalyticsCategoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_getCallAnalyticsCategoryCmd).Standalone()

		transcribe_getCallAnalyticsCategoryCmd.Flags().String("category-name", "", "The name of the Call Analytics category you want information about.")
		transcribe_getCallAnalyticsCategoryCmd.MarkFlagRequired("category-name")
	})
	transcribeCmd.AddCommand(transcribe_getCallAnalyticsCategoryCmd)
}
