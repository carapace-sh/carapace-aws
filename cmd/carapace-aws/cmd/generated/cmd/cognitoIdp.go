package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdpCmd = &cobra.Command{
	Use:   "cognito-idp",
	Short: "With the Amazon Cognito user pools API, you can configure user pools and authenticate users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdpCmd).Standalone()

	rootCmd.AddCommand(cognitoIdpCmd)
}
