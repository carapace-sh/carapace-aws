package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listUserTagsCmd = &cobra.Command{
	Use:   "list-user-tags",
	Short: "Lists the tags that are attached to the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listUserTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listUserTagsCmd).Standalone()

		iam_listUserTagsCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listUserTagsCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listUserTagsCmd.Flags().String("user-name", "", "The name of the IAM user whose tags you want to see.")
		iam_listUserTagsCmd.MarkFlagRequired("user-name")
	})
	iamCmd.AddCommand(iam_listUserTagsCmd)
}
