package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listRolesCmd = &cobra.Command{
	Use:   "list-roles",
	Short: "Lists the IAM roles that have the specified path prefix.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listRolesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listRolesCmd).Standalone()

		iam_listRolesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listRolesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listRolesCmd.Flags().String("path-prefix", "", "The path prefix for filtering the results.")
	})
	iamCmd.AddCommand(iam_listRolesCmd)
}
