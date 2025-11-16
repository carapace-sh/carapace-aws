package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listDocumentClassifierSummariesCmd = &cobra.Command{
	Use:   "list-document-classifier-summaries",
	Short: "Gets a list of summaries of the document classifiers that you have created",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listDocumentClassifierSummariesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_listDocumentClassifierSummariesCmd).Standalone()

		comprehend_listDocumentClassifierSummariesCmd.Flags().String("max-results", "", "The maximum number of results to return on each page.")
		comprehend_listDocumentClassifierSummariesCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	})
	comprehendCmd.AddCommand(comprehend_listDocumentClassifierSummariesCmd)
}
