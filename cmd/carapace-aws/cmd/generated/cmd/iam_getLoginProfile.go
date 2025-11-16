package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getLoginProfileCmd = &cobra.Command{
	Use:   "get-login-profile",
	Short: "Retrieves the user name for the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getLoginProfileCmd).Standalone()

	iam_getLoginProfileCmd.Flags().String("user-name", "", "The name of the user whose login profile you want to retrieve.")
	iamCmd.AddCommand(iam_getLoginProfileCmd)
}
