package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listGroupsForUserCmd = &cobra.Command{
	Use:   "list-groups-for-user",
	Short: "Lists the IAM groups that the specified IAM user belongs to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listGroupsForUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listGroupsForUserCmd).Standalone()

		iam_listGroupsForUserCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listGroupsForUserCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listGroupsForUserCmd.Flags().String("user-name", "", "The name of the user to list groups for.")
		iam_listGroupsForUserCmd.MarkFlagRequired("user-name")
	})
	iamCmd.AddCommand(iam_listGroupsForUserCmd)
}
