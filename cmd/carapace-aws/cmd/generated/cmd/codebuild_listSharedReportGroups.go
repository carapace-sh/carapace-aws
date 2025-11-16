package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listSharedReportGroupsCmd = &cobra.Command{
	Use:   "list-shared-report-groups",
	Short: "Gets a list of report groups that are shared with other Amazon Web Services accounts or users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listSharedReportGroupsCmd).Standalone()

	codebuild_listSharedReportGroupsCmd.Flags().String("max-results", "", "The maximum number of paginated shared report groups per response.")
	codebuild_listSharedReportGroupsCmd.Flags().String("next-token", "", "During a previous call, the maximum number of items that can be returned is the value specified in `maxResults`.")
	codebuild_listSharedReportGroupsCmd.Flags().String("sort-by", "", "The criterion to be used to list report groups shared with the current Amazon Web Services account or user.")
	codebuild_listSharedReportGroupsCmd.Flags().String("sort-order", "", "The order in which to list shared report groups.")
	codebuildCmd.AddCommand(codebuild_listSharedReportGroupsCmd)
}
