package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_listIdentitiesCmd = &cobra.Command{
	Use:   "list-identities",
	Short: "Lists the identities in an identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_listIdentitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdentity_listIdentitiesCmd).Standalone()

		cognitoIdentity_listIdentitiesCmd.Flags().String("hide-disabled", "", "An optional boolean parameter that allows you to hide disabled identities.")
		cognitoIdentity_listIdentitiesCmd.Flags().String("identity-pool-id", "", "An identity pool ID in the format REGION:GUID.")
		cognitoIdentity_listIdentitiesCmd.Flags().String("max-results", "", "The maximum number of identities to return.")
		cognitoIdentity_listIdentitiesCmd.Flags().String("next-token", "", "A pagination token.")
		cognitoIdentity_listIdentitiesCmd.MarkFlagRequired("identity-pool-id")
		cognitoIdentity_listIdentitiesCmd.MarkFlagRequired("max-results")
	})
	cognitoIdentityCmd.AddCommand(cognitoIdentity_listIdentitiesCmd)
}
