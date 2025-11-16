package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "Lists the IAM groups that have the specified path prefix.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listGroupsCmd).Standalone()

	iam_listGroupsCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_listGroupsCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iam_listGroupsCmd.Flags().String("path-prefix", "", "The path prefix for filtering the results.")
	iamCmd.AddCommand(iam_listGroupsCmd)
}
