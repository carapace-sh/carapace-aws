package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_getPrincipalTagAttributeMapCmd = &cobra.Command{
	Use:   "get-principal-tag-attribute-map",
	Short: "Use `GetPrincipalTagAttributeMap` to list all mappings between `PrincipalTags` and user attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_getPrincipalTagAttributeMapCmd).Standalone()

	cognitoIdentity_getPrincipalTagAttributeMapCmd.Flags().String("identity-pool-id", "", "You can use this operation to get the ID of the Identity Pool you setup attribute mappings for.")
	cognitoIdentity_getPrincipalTagAttributeMapCmd.Flags().String("identity-provider-name", "", "You can use this operation to get the provider name.")
	cognitoIdentity_getPrincipalTagAttributeMapCmd.MarkFlagRequired("identity-pool-id")
	cognitoIdentity_getPrincipalTagAttributeMapCmd.MarkFlagRequired("identity-provider-name")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_getPrincipalTagAttributeMapCmd)
}
