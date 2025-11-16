package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_deleteTermsCmd = &cobra.Command{
	Use:   "delete-terms",
	Short: "Deletes the terms documents with the requested ID from your app client.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_deleteTermsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_deleteTermsCmd).Standalone()

		cognitoIdp_deleteTermsCmd.Flags().String("terms-id", "", "The ID of the terms documents that you want to delete.")
		cognitoIdp_deleteTermsCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the terms documents that you want to delete.")
		cognitoIdp_deleteTermsCmd.MarkFlagRequired("terms-id")
		cognitoIdp_deleteTermsCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_deleteTermsCmd)
}
