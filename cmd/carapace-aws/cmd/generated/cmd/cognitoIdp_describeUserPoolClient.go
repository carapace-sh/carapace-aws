package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_describeUserPoolClientCmd = &cobra.Command{
	Use:   "describe-user-pool-client",
	Short: "Given an app client ID, returns configuration information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_describeUserPoolClientCmd).Standalone()

	cognitoIdp_describeUserPoolClientCmd.Flags().String("client-id", "", "The ID of the app client that you want to describe.")
	cognitoIdp_describeUserPoolClientCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the app client you want to describe.")
	cognitoIdp_describeUserPoolClientCmd.MarkFlagRequired("client-id")
	cognitoIdp_describeUserPoolClientCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_describeUserPoolClientCmd)
}
