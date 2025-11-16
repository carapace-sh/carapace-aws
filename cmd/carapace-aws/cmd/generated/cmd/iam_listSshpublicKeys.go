package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listSshpublicKeysCmd = &cobra.Command{
	Use:   "list-sshpublic-keys",
	Short: "Returns information about the SSH public keys associated with the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listSshpublicKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listSshpublicKeysCmd).Standalone()

		iam_listSshpublicKeysCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listSshpublicKeysCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listSshpublicKeysCmd.Flags().String("user-name", "", "The name of the IAM user to list SSH public keys for.")
	})
	iamCmd.AddCommand(iam_listSshpublicKeysCmd)
}
