package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listSamlprovidersCmd = &cobra.Command{
	Use:   "list-samlproviders",
	Short: "Lists the SAML provider resource objects defined in IAM in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listSamlprovidersCmd).Standalone()

	iamCmd.AddCommand(iam_listSamlprovidersCmd)
}
