package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listMfadevicesCmd = &cobra.Command{
	Use:   "list-mfadevices",
	Short: "Lists the MFA devices for an IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listMfadevicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listMfadevicesCmd).Standalone()

		iam_listMfadevicesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listMfadevicesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listMfadevicesCmd.Flags().String("user-name", "", "The name of the user whose MFA devices you want to list.")
	})
	iamCmd.AddCommand(iam_listMfadevicesCmd)
}
