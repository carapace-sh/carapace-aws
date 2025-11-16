package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_listKeywordsForDataSourceCmd = &cobra.Command{
	Use:   "list-keywords-for-data-source",
	Short: "Returns a list of keywords that are pre-mapped to the specified control data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_listKeywordsForDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_listKeywordsForDataSourceCmd).Standalone()

		auditmanager_listKeywordsForDataSourceCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
		auditmanager_listKeywordsForDataSourceCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		auditmanager_listKeywordsForDataSourceCmd.Flags().String("source", "", "The control mapping data source that the keywords apply to.")
		auditmanager_listKeywordsForDataSourceCmd.MarkFlagRequired("source")
	})
	auditmanagerCmd.AddCommand(auditmanager_listKeywordsForDataSourceCmd)
}
