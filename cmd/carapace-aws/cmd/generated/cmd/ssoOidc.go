package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoOidcCmd = &cobra.Command{
	Use:   "sso-oidc",
	Short: "IAM Identity Center OpenID Connect (OIDC) is a web service that enables a client (such as CLI or a native application) to register with IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoOidcCmd).Standalone()

	rootCmd.AddCommand(ssoOidcCmd)
}
