package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_listUserPoolsCmd = &cobra.Command{
	Use:   "list-user-pools",
	Short: "Lists user pools and their details in the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_listUserPoolsCmd).Standalone()

	cognitoIdp_listUserPoolsCmd.Flags().String("max-results", "", "The maximum number of user pools that you want Amazon Cognito to return in the response.")
	cognitoIdp_listUserPoolsCmd.Flags().String("next-token", "", "This API operation returns a limited number of results.")
	cognitoIdp_listUserPoolsCmd.MarkFlagRequired("max-results")
	cognitoIdpCmd.AddCommand(cognitoIdp_listUserPoolsCmd)
}
