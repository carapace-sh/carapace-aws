package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getGroupCmd = &cobra.Command{
	Use:   "get-group",
	Short: "Returns a list of IAM users that are in the specified IAM group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getGroupCmd).Standalone()

	iam_getGroupCmd.Flags().String("group-name", "", "The name of the group.")
	iam_getGroupCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_getGroupCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iam_getGroupCmd.MarkFlagRequired("group-name")
	iamCmd.AddCommand(iam_getGroupCmd)
}
