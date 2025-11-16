package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_listImportJobsCmd = &cobra.Command{
	Use:   "list-import-jobs",
	Short: "Lists information about import jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_listImportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_listImportJobsCmd).Standalone()

		wisdom_listImportJobsCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		wisdom_listImportJobsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		wisdom_listImportJobsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		wisdom_listImportJobsCmd.MarkFlagRequired("knowledge-base-id")
	})
	wisdomCmd.AddCommand(wisdom_listImportJobsCmd)
}
