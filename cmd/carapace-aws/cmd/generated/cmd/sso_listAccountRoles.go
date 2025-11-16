package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sso_listAccountRolesCmd = &cobra.Command{
	Use:   "list-account-roles",
	Short: "Lists all roles that are assigned to the user for a given AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sso_listAccountRolesCmd).Standalone()

	sso_listAccountRolesCmd.Flags().String("access-token", "", "The token issued by the `CreateToken` API call.")
	sso_listAccountRolesCmd.Flags().String("account-id", "", "The identifier for the AWS account that is assigned to the user.")
	sso_listAccountRolesCmd.Flags().String("max-results", "", "The number of items that clients can request per page.")
	sso_listAccountRolesCmd.Flags().String("next-token", "", "The page token from the previous response output when you request subsequent pages.")
	sso_listAccountRolesCmd.MarkFlagRequired("access-token")
	sso_listAccountRolesCmd.MarkFlagRequired("account-id")
	ssoCmd.AddCommand(sso_listAccountRolesCmd)
}
