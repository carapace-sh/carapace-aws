package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_createAccountCmd = &cobra.Command{
	Use:   "create-account",
	Short: "Creates an Amazon Web Services account that is automatically a member of the organization whose credentials made the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_createAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_createAccountCmd).Standalone()

		organizations_createAccountCmd.Flags().String("account-name", "", "The friendly name of the member account.")
		organizations_createAccountCmd.Flags().String("email", "", "The email address of the owner to assign to the new member account.")
		organizations_createAccountCmd.Flags().String("iam-user-access-to-billing", "", "If set to `ALLOW`, the new account enables IAM users to access account billing information *if* they have the required permissions.")
		organizations_createAccountCmd.Flags().String("role-name", "", "The name of an IAM role that Organizations automatically preconfigures in the new member account.")
		organizations_createAccountCmd.Flags().String("tags", "", "A list of tags that you want to attach to the newly created account.")
		organizations_createAccountCmd.MarkFlagRequired("account-name")
		organizations_createAccountCmd.MarkFlagRequired("email")
	})
	organizationsCmd.AddCommand(organizations_createAccountCmd)
}
