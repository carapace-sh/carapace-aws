package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_listUserPoolClientsCmd = &cobra.Command{
	Use:   "list-user-pool-clients",
	Short: "Given a user pool ID, lists app clients.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_listUserPoolClientsCmd).Standalone()

	cognitoIdp_listUserPoolClientsCmd.Flags().String("max-results", "", "The maximum number of app clients that you want Amazon Cognito to return in the response.")
	cognitoIdp_listUserPoolClientsCmd.Flags().String("next-token", "", "This API operation returns a limited number of results.")
	cognitoIdp_listUserPoolClientsCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to list user pool clients.")
	cognitoIdp_listUserPoolClientsCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_listUserPoolClientsCmd)
}
