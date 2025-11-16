package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listImportJobsCmd = &cobra.Command{
	Use:   "list-import-jobs",
	Short: "Lists information about import jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listImportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_listImportJobsCmd).Standalone()

		qconnect_listImportJobsCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_listImportJobsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		qconnect_listImportJobsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		qconnect_listImportJobsCmd.MarkFlagRequired("knowledge-base-id")
	})
	qconnectCmd.AddCommand(qconnect_listImportJobsCmd)
}
