package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags for a top-level EFS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_listTagsForResourceCmd).Standalone()

		efs_listTagsForResourceCmd.Flags().String("max-results", "", "(Optional) Specifies the maximum number of tag objects to return in the response.")
		efs_listTagsForResourceCmd.Flags().String("next-token", "", "(Optional) You can use `NextToken` in a subsequent request to fetch the next page of access point descriptions if the response payload was paginated.")
		efs_listTagsForResourceCmd.Flags().String("resource-id", "", "Specifies the EFS resource you want to retrieve tags for.")
		efs_listTagsForResourceCmd.MarkFlagRequired("resource-id")
	})
	efsCmd.AddCommand(efs_listTagsForResourceCmd)
}
