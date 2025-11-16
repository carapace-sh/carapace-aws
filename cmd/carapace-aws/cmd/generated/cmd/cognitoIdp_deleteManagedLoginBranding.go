package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_deleteManagedLoginBrandingCmd = &cobra.Command{
	Use:   "delete-managed-login-branding",
	Short: "Deletes a managed login branding style.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_deleteManagedLoginBrandingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_deleteManagedLoginBrandingCmd).Standalone()

		cognitoIdp_deleteManagedLoginBrandingCmd.Flags().String("managed-login-branding-id", "", "The ID of the managed login branding style that you want to delete.")
		cognitoIdp_deleteManagedLoginBrandingCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the managed login branding style that you want to delete.")
		cognitoIdp_deleteManagedLoginBrandingCmd.MarkFlagRequired("managed-login-branding-id")
		cognitoIdp_deleteManagedLoginBrandingCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_deleteManagedLoginBrandingCmd)
}
