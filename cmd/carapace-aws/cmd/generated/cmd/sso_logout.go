package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sso_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Removes the locally stored SSO tokens from the client-side cache and sends an API call to the IAM Identity Center service to invalidate the corresponding server-side IAM Identity Center sign in session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sso_logoutCmd).Standalone()

	sso_logoutCmd.Flags().String("access-token", "", "The token issued by the `CreateToken` API call.")
	sso_logoutCmd.MarkFlagRequired("access-token")
	ssoCmd.AddCommand(sso_logoutCmd)
}
