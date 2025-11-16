package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getRoleCmd = &cobra.Command{
	Use:   "get-role",
	Short: "Retrieves information about the specified role, including the role's path, GUID, ARN, and the role's trust policy that grants permission to assume the role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getRoleCmd).Standalone()

	iam_getRoleCmd.Flags().String("role-name", "", "The name of the IAM role to get information about.")
	iam_getRoleCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_getRoleCmd)
}
