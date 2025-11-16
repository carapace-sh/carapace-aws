package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listUserPoliciesCmd = &cobra.Command{
	Use:   "list-user-policies",
	Short: "Lists the names of the inline policies embedded in the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listUserPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listUserPoliciesCmd).Standalone()

		iam_listUserPoliciesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listUserPoliciesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listUserPoliciesCmd.Flags().String("user-name", "", "The name of the user to list policies for.")
		iam_listUserPoliciesCmd.MarkFlagRequired("user-name")
	})
	iamCmd.AddCommand(iam_listUserPoliciesCmd)
}
