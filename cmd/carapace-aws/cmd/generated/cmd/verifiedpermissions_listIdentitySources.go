package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_listIdentitySourcesCmd = &cobra.Command{
	Use:   "list-identity-sources",
	Short: "Returns a paginated list of all of the identity sources defined in the specified policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_listIdentitySourcesCmd).Standalone()

	verifiedpermissions_listIdentitySourcesCmd.Flags().String("filters", "", "Specifies characteristics of an identity source that you can use to limit the output to matching identity sources.")
	verifiedpermissions_listIdentitySourcesCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included in each response.")
	verifiedpermissions_listIdentitySourcesCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	verifiedpermissions_listIdentitySourcesCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that contains the identity sources that you want to list.")
	verifiedpermissions_listIdentitySourcesCmd.MarkFlagRequired("policy-store-id")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_listIdentitySourcesCmd)
}
