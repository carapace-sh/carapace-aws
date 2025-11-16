package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a user who can be used in WorkMail by calling the [RegisterToWorkMail]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_createUserCmd).Standalone()

	workmail_createUserCmd.Flags().String("display-name", "", "The display name for the new user.")
	workmail_createUserCmd.Flags().String("first-name", "", "The first name of the new user.")
	workmail_createUserCmd.Flags().Bool("hidden-from-global-address-list", false, "If this parameter is enabled, the user will be hidden from the address book.")
	workmail_createUserCmd.Flags().String("identity-provider-user-id", "", "User ID from the IAM Identity Center.")
	workmail_createUserCmd.Flags().String("last-name", "", "The last name of the new user.")
	workmail_createUserCmd.Flags().String("name", "", "The name for the new user.")
	workmail_createUserCmd.Flags().Bool("no-hidden-from-global-address-list", false, "If this parameter is enabled, the user will be hidden from the address book.")
	workmail_createUserCmd.Flags().String("organization-id", "", "The identifier of the organization for which the user is created.")
	workmail_createUserCmd.Flags().String("password", "", "The password for the new user.")
	workmail_createUserCmd.Flags().String("role", "", "The role of the new user.")
	workmail_createUserCmd.MarkFlagRequired("display-name")
	workmail_createUserCmd.MarkFlagRequired("name")
	workmail_createUserCmd.Flag("no-hidden-from-global-address-list").Hidden = true
	workmail_createUserCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_createUserCmd)
}
