package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getUserCmd = &cobra.Command{
	Use:   "get-user",
	Short: "Retrieves information about the specified IAM user, including the user's creation date, path, unique ID, and ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getUserCmd).Standalone()

	iam_getUserCmd.Flags().String("user-name", "", "The name of the user to get information about.")
	iamCmd.AddCommand(iam_getUserCmd)
}
