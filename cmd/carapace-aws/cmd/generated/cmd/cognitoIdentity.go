package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentityCmd = &cobra.Command{
	Use:   "cognito-identity",
	Short: "Amazon Cognito Federated Identities\n\nAmazon Cognito Federated Identities is a web service that delivers scoped temporary credentials to mobile devices and other untrusted environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdentityCmd).Standalone()

	})
	rootCmd.AddCommand(cognitoIdentityCmd)
}
