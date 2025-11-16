package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_createManagedLoginBrandingCmd = &cobra.Command{
	Use:   "create-managed-login-branding",
	Short: "Creates a new set of branding settings for a user pool style and associates it with an app client.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_createManagedLoginBrandingCmd).Standalone()

	cognitoIdp_createManagedLoginBrandingCmd.Flags().String("assets", "", "An array of image files that you want to apply to functions like backgrounds, logos, and icons.")
	cognitoIdp_createManagedLoginBrandingCmd.Flags().String("client-id", "", "The app client that you want to create the branding style for.")
	cognitoIdp_createManagedLoginBrandingCmd.Flags().String("settings", "", "A JSON file, encoded as a `Document` type, with the the settings that you want to apply to your style.")
	cognitoIdp_createManagedLoginBrandingCmd.Flags().String("use-cognito-provided-values", "", "When true, applies the default branding style options.")
	cognitoIdp_createManagedLoginBrandingCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to create a new branding style.")
	cognitoIdp_createManagedLoginBrandingCmd.MarkFlagRequired("client-id")
	cognitoIdp_createManagedLoginBrandingCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_createManagedLoginBrandingCmd)
}
