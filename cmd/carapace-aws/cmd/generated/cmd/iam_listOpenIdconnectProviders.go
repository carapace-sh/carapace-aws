package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listOpenIdconnectProvidersCmd = &cobra.Command{
	Use:   "list-open-idconnect-providers",
	Short: "Lists information about the IAM OpenID Connect (OIDC) provider resource objects defined in the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listOpenIdconnectProvidersCmd).Standalone()

	iamCmd.AddCommand(iam_listOpenIdconnectProvidersCmd)
}
