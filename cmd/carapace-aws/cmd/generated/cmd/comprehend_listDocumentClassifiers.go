package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listDocumentClassifiersCmd = &cobra.Command{
	Use:   "list-document-classifiers",
	Short: "Gets a list of the document classifiers that you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listDocumentClassifiersCmd).Standalone()

	comprehend_listDocumentClassifiersCmd.Flags().String("filter", "", "Filters the jobs that are returned.")
	comprehend_listDocumentClassifiersCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
	comprehend_listDocumentClassifiersCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	comprehendCmd.AddCommand(comprehend_listDocumentClassifiersCmd)
}
