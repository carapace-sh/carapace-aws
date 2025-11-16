package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_unlinkDeveloperIdentityCmd = &cobra.Command{
	Use:   "unlink-developer-identity",
	Short: "Unlinks a `DeveloperUserIdentifier` from an existing identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_unlinkDeveloperIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdentity_unlinkDeveloperIdentityCmd).Standalone()

		cognitoIdentity_unlinkDeveloperIdentityCmd.Flags().String("developer-provider-name", "", "The \"domain\" by which Cognito will refer to your users.")
		cognitoIdentity_unlinkDeveloperIdentityCmd.Flags().String("developer-user-identifier", "", "A unique ID used by your backend authentication process to identify a user.")
		cognitoIdentity_unlinkDeveloperIdentityCmd.Flags().String("identity-id", "", "A unique identifier in the format REGION:GUID.")
		cognitoIdentity_unlinkDeveloperIdentityCmd.Flags().String("identity-pool-id", "", "An identity pool ID in the format REGION:GUID.")
		cognitoIdentity_unlinkDeveloperIdentityCmd.MarkFlagRequired("developer-provider-name")
		cognitoIdentity_unlinkDeveloperIdentityCmd.MarkFlagRequired("developer-user-identifier")
		cognitoIdentity_unlinkDeveloperIdentityCmd.MarkFlagRequired("identity-id")
		cognitoIdentity_unlinkDeveloperIdentityCmd.MarkFlagRequired("identity-pool-id")
	})
	cognitoIdentityCmd.AddCommand(cognitoIdentity_unlinkDeveloperIdentityCmd)
}
