package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_deleteIdentityPoolCmd = &cobra.Command{
	Use:   "delete-identity-pool",
	Short: "Deletes an identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_deleteIdentityPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdentity_deleteIdentityPoolCmd).Standalone()

		cognitoIdentity_deleteIdentityPoolCmd.Flags().String("identity-pool-id", "", "An identity pool ID in the format REGION:GUID.")
		cognitoIdentity_deleteIdentityPoolCmd.MarkFlagRequired("identity-pool-id")
	})
	cognitoIdentityCmd.AddCommand(cognitoIdentity_deleteIdentityPoolCmd)
}
