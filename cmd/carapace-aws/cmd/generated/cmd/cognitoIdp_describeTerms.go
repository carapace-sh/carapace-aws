package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_describeTermsCmd = &cobra.Command{
	Use:   "describe-terms",
	Short: "Returns details for the requested terms documents ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_describeTermsCmd).Standalone()

	cognitoIdp_describeTermsCmd.Flags().String("terms-id", "", "The ID of the terms documents that you want to describe.")
	cognitoIdp_describeTermsCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the terms documents that you want to describe.")
	cognitoIdp_describeTermsCmd.MarkFlagRequired("terms-id")
	cognitoIdp_describeTermsCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_describeTermsCmd)
}
