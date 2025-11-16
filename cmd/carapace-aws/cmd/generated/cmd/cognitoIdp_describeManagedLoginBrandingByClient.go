package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_describeManagedLoginBrandingByClientCmd = &cobra.Command{
	Use:   "describe-managed-login-branding-by-client",
	Short: "Given the ID of a user pool app client, returns detailed information about the style assigned to the app client.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_describeManagedLoginBrandingByClientCmd).Standalone()

	cognitoIdp_describeManagedLoginBrandingByClientCmd.Flags().String("client-id", "", "The app client that's assigned to the branding style that you want more information about.")
	cognitoIdp_describeManagedLoginBrandingByClientCmd.Flags().String("return-merged-resources", "", "When `true`, returns values for branding options that are unchanged from Amazon Cognito defaults.")
	cognitoIdp_describeManagedLoginBrandingByClientCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the app client where you want more information about the managed login branding style.")
	cognitoIdp_describeManagedLoginBrandingByClientCmd.MarkFlagRequired("client-id")
	cognitoIdp_describeManagedLoginBrandingByClientCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_describeManagedLoginBrandingByClientCmd)
}
