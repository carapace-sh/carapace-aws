package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_updateManagedLoginBrandingCmd = &cobra.Command{
	Use:   "update-managed-login-branding",
	Short: "Configures the branding settings for a user pool style.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_updateManagedLoginBrandingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_updateManagedLoginBrandingCmd).Standalone()

		cognitoIdp_updateManagedLoginBrandingCmd.Flags().String("assets", "", "An array of image files that you want to apply to roles like backgrounds, logos, and icons.")
		cognitoIdp_updateManagedLoginBrandingCmd.Flags().String("managed-login-branding-id", "", "The ID of the managed login branding style that you want to update.")
		cognitoIdp_updateManagedLoginBrandingCmd.Flags().String("settings", "", "A JSON file, encoded as a `Document` type, with the the settings that you want to apply to your style.")
		cognitoIdp_updateManagedLoginBrandingCmd.Flags().String("use-cognito-provided-values", "", "When `true`, applies the default branding style options.")
		cognitoIdp_updateManagedLoginBrandingCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the managed login branding style that you want to update.")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_updateManagedLoginBrandingCmd)
}
