package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentityCmd = &cobra.Command{
	Use:   "cognito-identity",
	Short: "Amazon Cognito Federated Identities",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdentityCmd).Standalone()

	})
	rootCmd.AddCommand(cognitoIdentityCmd)
}
