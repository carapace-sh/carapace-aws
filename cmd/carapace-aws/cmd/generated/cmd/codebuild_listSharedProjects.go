package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listSharedProjectsCmd = &cobra.Command{
	Use:   "list-shared-projects",
	Short: "Gets a list of projects that are shared with other Amazon Web Services accounts or users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listSharedProjectsCmd).Standalone()

	codebuild_listSharedProjectsCmd.Flags().String("max-results", "", "The maximum number of paginated shared build projects returned per response.")
	codebuild_listSharedProjectsCmd.Flags().String("next-token", "", "During a previous call, the maximum number of items that can be returned is the value specified in `maxResults`.")
	codebuild_listSharedProjectsCmd.Flags().String("sort-by", "", "The criterion to be used to list build projects shared with the current Amazon Web Services account or user.")
	codebuild_listSharedProjectsCmd.Flags().String("sort-order", "", "The order in which to list shared build projects.")
	codebuildCmd.AddCommand(codebuild_listSharedProjectsCmd)
}
