package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createAccessKeyCmd = &cobra.Command{
	Use:   "create-access-key",
	Short: "Creates a new Amazon Web Services secret access key and corresponding Amazon Web Services access key ID for the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createAccessKeyCmd).Standalone()

	iam_createAccessKeyCmd.Flags().String("user-name", "", "The name of the IAM user that the new key will belong to.")
	iamCmd.AddCommand(iam_createAccessKeyCmd)
}
