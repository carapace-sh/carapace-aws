package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_batchIsAuthorizedWithTokenCmd = &cobra.Command{
	Use:   "batch-is-authorized-with-token",
	Short: "Makes a series of decisions about multiple authorization requests for one token.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_batchIsAuthorizedWithTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_batchIsAuthorizedWithTokenCmd).Standalone()

		verifiedpermissions_batchIsAuthorizedWithTokenCmd.Flags().String("access-token", "", "Specifies an access token for the principal that you want to authorize in each request.")
		verifiedpermissions_batchIsAuthorizedWithTokenCmd.Flags().String("entities", "", "(Optional) Specifies the list of resources and their associated attributes that Verified Permissions can examine when evaluating the policies.")
		verifiedpermissions_batchIsAuthorizedWithTokenCmd.Flags().String("identity-token", "", "Specifies an identity (ID) token for the principal that you want to authorize in each request.")
		verifiedpermissions_batchIsAuthorizedWithTokenCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store.")
		verifiedpermissions_batchIsAuthorizedWithTokenCmd.Flags().String("requests", "", "An array of up to 30 requests that you want Verified Permissions to evaluate.")
		verifiedpermissions_batchIsAuthorizedWithTokenCmd.MarkFlagRequired("policy-store-id")
		verifiedpermissions_batchIsAuthorizedWithTokenCmd.MarkFlagRequired("requests")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_batchIsAuthorizedWithTokenCmd)
}
