package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_listAdaptersCmd = &cobra.Command{
	Use:   "list-adapters",
	Short: "Lists all adapters that match the specified filtration criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_listAdaptersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(textract_listAdaptersCmd).Standalone()

		textract_listAdaptersCmd.Flags().String("after-creation-time", "", "Specifies the lower bound for the ListAdapters operation.")
		textract_listAdaptersCmd.Flags().String("before-creation-time", "", "Specifies the upper bound for the ListAdapters operation.")
		textract_listAdaptersCmd.Flags().String("max-results", "", "The maximum number of results to return when listing adapters.")
		textract_listAdaptersCmd.Flags().String("next-token", "", "Identifies the next page of results to return when listing adapters.")
	})
	textractCmd.AddCommand(textract_listAdaptersCmd)
}
