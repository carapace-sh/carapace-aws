package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_mergeDeveloperIdentitiesCmd = &cobra.Command{
	Use:   "merge-developer-identities",
	Short: "Merges two users having different `IdentityId`s, existing in the same identity pool, and identified by the same developer provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_mergeDeveloperIdentitiesCmd).Standalone()

	cognitoIdentity_mergeDeveloperIdentitiesCmd.Flags().String("destination-user-identifier", "", "User identifier for the destination user.")
	cognitoIdentity_mergeDeveloperIdentitiesCmd.Flags().String("developer-provider-name", "", "The \"domain\" by which Cognito will refer to your users.")
	cognitoIdentity_mergeDeveloperIdentitiesCmd.Flags().String("identity-pool-id", "", "An identity pool ID in the format REGION:GUID.")
	cognitoIdentity_mergeDeveloperIdentitiesCmd.Flags().String("source-user-identifier", "", "User identifier for the source user.")
	cognitoIdentity_mergeDeveloperIdentitiesCmd.MarkFlagRequired("destination-user-identifier")
	cognitoIdentity_mergeDeveloperIdentitiesCmd.MarkFlagRequired("developer-provider-name")
	cognitoIdentity_mergeDeveloperIdentitiesCmd.MarkFlagRequired("identity-pool-id")
	cognitoIdentity_mergeDeveloperIdentitiesCmd.MarkFlagRequired("source-user-identifier")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_mergeDeveloperIdentitiesCmd)
}
