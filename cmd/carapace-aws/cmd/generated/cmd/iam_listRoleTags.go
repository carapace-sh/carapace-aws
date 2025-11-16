package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listRoleTagsCmd = &cobra.Command{
	Use:   "list-role-tags",
	Short: "Lists the tags that are attached to the specified role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listRoleTagsCmd).Standalone()

	iam_listRoleTagsCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_listRoleTagsCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iam_listRoleTagsCmd.Flags().String("role-name", "", "The name of the IAM role for which you want to see the list of tags.")
	iam_listRoleTagsCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_listRoleTagsCmd)
}
