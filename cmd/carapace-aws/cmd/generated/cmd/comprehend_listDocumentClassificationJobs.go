package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listDocumentClassificationJobsCmd = &cobra.Command{
	Use:   "list-document-classification-jobs",
	Short: "Gets a list of the documentation classification jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listDocumentClassificationJobsCmd).Standalone()

	comprehend_listDocumentClassificationJobsCmd.Flags().String("filter", "", "Filters the jobs that are returned.")
	comprehend_listDocumentClassificationJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
	comprehend_listDocumentClassificationJobsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	comprehendCmd.AddCommand(comprehend_listDocumentClassificationJobsCmd)
}
