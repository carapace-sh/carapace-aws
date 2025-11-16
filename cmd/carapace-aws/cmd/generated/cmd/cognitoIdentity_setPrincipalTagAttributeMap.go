package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_setPrincipalTagAttributeMapCmd = &cobra.Command{
	Use:   "set-principal-tag-attribute-map",
	Short: "You can use this operation to use default (username and clientID) attribute or custom attribute mappings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_setPrincipalTagAttributeMapCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdentity_setPrincipalTagAttributeMapCmd).Standalone()

		cognitoIdentity_setPrincipalTagAttributeMapCmd.Flags().String("identity-pool-id", "", "The ID of the Identity Pool you want to set attribute mappings for.")
		cognitoIdentity_setPrincipalTagAttributeMapCmd.Flags().String("identity-provider-name", "", "The provider name you want to use for attribute mappings.")
		cognitoIdentity_setPrincipalTagAttributeMapCmd.Flags().String("principal-tags", "", "You can use this operation to add principal tags.")
		cognitoIdentity_setPrincipalTagAttributeMapCmd.Flags().String("use-defaults", "", "You can use this operation to use default (username and clientID) attribute mappings.")
		cognitoIdentity_setPrincipalTagAttributeMapCmd.MarkFlagRequired("identity-pool-id")
		cognitoIdentity_setPrincipalTagAttributeMapCmd.MarkFlagRequired("identity-provider-name")
	})
	cognitoIdentityCmd.AddCommand(cognitoIdentity_setPrincipalTagAttributeMapCmd)
}
