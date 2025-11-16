package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listReportGroupsCmd = &cobra.Command{
	Use:   "list-report-groups",
	Short: "Gets a list ARNs for the report groups in the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listReportGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_listReportGroupsCmd).Standalone()

		codebuild_listReportGroupsCmd.Flags().String("max-results", "", "The maximum number of paginated report groups returned per response.")
		codebuild_listReportGroupsCmd.Flags().String("next-token", "", "During a previous call, the maximum number of items that can be returned is the value specified in `maxResults`.")
		codebuild_listReportGroupsCmd.Flags().String("sort-by", "", "The criterion to be used to list build report groups.")
		codebuild_listReportGroupsCmd.Flags().String("sort-order", "", "Used to specify the order to sort the list of returned report groups.")
	})
	codebuildCmd.AddCommand(codebuild_listReportGroupsCmd)
}
