package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_deleteCallAnalyticsCategoryCmd = &cobra.Command{
	Use:   "delete-call-analytics-category",
	Short: "Deletes a Call Analytics category.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_deleteCallAnalyticsCategoryCmd).Standalone()

	transcribe_deleteCallAnalyticsCategoryCmd.Flags().String("category-name", "", "The name of the Call Analytics category you want to delete.")
	transcribe_deleteCallAnalyticsCategoryCmd.MarkFlagRequired("category-name")
	transcribeCmd.AddCommand(transcribe_deleteCallAnalyticsCategoryCmd)
}
