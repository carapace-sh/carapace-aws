package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_updateTermsCmd = &cobra.Command{
	Use:   "update-terms",
	Short: "Modifies existing terms documents for the requested app client.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_updateTermsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_updateTermsCmd).Standalone()

		cognitoIdp_updateTermsCmd.Flags().String("enforcement", "", "This parameter is reserved for future use and currently accepts only one value.")
		cognitoIdp_updateTermsCmd.Flags().String("links", "", "A map of URLs to languages.")
		cognitoIdp_updateTermsCmd.Flags().String("terms-id", "", "The ID of the terms document that you want to update.")
		cognitoIdp_updateTermsCmd.Flags().String("terms-name", "", "The new name that you want to apply to the requested terms documents.")
		cognitoIdp_updateTermsCmd.Flags().String("terms-source", "", "This parameter is reserved for future use and currently accepts only one value.")
		cognitoIdp_updateTermsCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the terms that you want to update.")
		cognitoIdp_updateTermsCmd.MarkFlagRequired("terms-id")
		cognitoIdp_updateTermsCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_updateTermsCmd)
}
