package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_listTermsCmd = &cobra.Command{
	Use:   "list-terms",
	Short: "Returns details about all terms documents for the requested user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_listTermsCmd).Standalone()

	cognitoIdp_listTermsCmd.Flags().String("max-results", "", "The maximum number of terms documents that you want Amazon Cognito to return in the response.")
	cognitoIdp_listTermsCmd.Flags().String("next-token", "", "This API operation returns a limited number of results.")
	cognitoIdp_listTermsCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to list terms documents.")
	cognitoIdp_listTermsCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_listTermsCmd)
}
