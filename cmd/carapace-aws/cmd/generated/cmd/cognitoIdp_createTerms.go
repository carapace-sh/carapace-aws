package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_createTermsCmd = &cobra.Command{
	Use:   "create-terms",
	Short: "Creates terms documents for the requested app client.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_createTermsCmd).Standalone()

	cognitoIdp_createTermsCmd.Flags().String("client-id", "", "The ID of the app client where you want to create terms documents.")
	cognitoIdp_createTermsCmd.Flags().String("enforcement", "", "This parameter is reserved for future use and currently accepts only one value.")
	cognitoIdp_createTermsCmd.Flags().String("links", "", "A map of URLs to languages.")
	cognitoIdp_createTermsCmd.Flags().String("terms-name", "", "A friendly name for the document that you want to create in the current request.")
	cognitoIdp_createTermsCmd.Flags().String("terms-source", "", "This parameter is reserved for future use and currently accepts only one value.")
	cognitoIdp_createTermsCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to create terms documents.")
	cognitoIdp_createTermsCmd.MarkFlagRequired("client-id")
	cognitoIdp_createTermsCmd.MarkFlagRequired("enforcement")
	cognitoIdp_createTermsCmd.MarkFlagRequired("terms-name")
	cognitoIdp_createTermsCmd.MarkFlagRequired("terms-source")
	cognitoIdp_createTermsCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_createTermsCmd)
}
