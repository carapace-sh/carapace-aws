package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listAccessKeysCmd = &cobra.Command{
	Use:   "list-access-keys",
	Short: "Returns information about the access key IDs associated with the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listAccessKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listAccessKeysCmd).Standalone()

		iam_listAccessKeysCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listAccessKeysCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listAccessKeysCmd.Flags().String("user-name", "", "The name of the user.")
	})
	iamCmd.AddCommand(iam_listAccessKeysCmd)
}
