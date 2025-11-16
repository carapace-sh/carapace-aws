package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_describeManagedLoginBrandingCmd = &cobra.Command{
	Use:   "describe-managed-login-branding",
	Short: "Given the ID of a managed login branding style, returns detailed information about the style.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_describeManagedLoginBrandingCmd).Standalone()

	cognitoIdp_describeManagedLoginBrandingCmd.Flags().String("managed-login-branding-id", "", "The ID of the managed login branding style that you want to get more information about.")
	cognitoIdp_describeManagedLoginBrandingCmd.Flags().String("return-merged-resources", "", "When `true`, returns values for branding options that are unchanged from Amazon Cognito defaults.")
	cognitoIdp_describeManagedLoginBrandingCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the managed login branding style that you want to get information about.")
	cognitoIdp_describeManagedLoginBrandingCmd.MarkFlagRequired("managed-login-branding-id")
	cognitoIdp_describeManagedLoginBrandingCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_describeManagedLoginBrandingCmd)
}
