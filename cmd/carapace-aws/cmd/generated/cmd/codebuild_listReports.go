package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listReportsCmd = &cobra.Command{
	Use:   "list-reports",
	Short: "Returns a list of ARNs for the reports in the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listReportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_listReportsCmd).Standalone()

		codebuild_listReportsCmd.Flags().String("filter", "", "A `ReportFilter` object used to filter the returned reports.")
		codebuild_listReportsCmd.Flags().String("max-results", "", "The maximum number of paginated reports returned per response.")
		codebuild_listReportsCmd.Flags().String("next-token", "", "During a previous call, the maximum number of items that can be returned is the value specified in `maxResults`.")
		codebuild_listReportsCmd.Flags().String("sort-order", "", "Specifies the sort order for the list of returned reports.")
	})
	codebuildCmd.AddCommand(codebuild_listReportsCmd)
}
