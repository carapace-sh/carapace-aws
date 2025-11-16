package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_batchIsAuthorizedCmd = &cobra.Command{
	Use:   "batch-is-authorized",
	Short: "Makes a series of decisions about multiple authorization requests for one principal or resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_batchIsAuthorizedCmd).Standalone()

	verifiedpermissions_batchIsAuthorizedCmd.Flags().String("entities", "", "(Optional) Specifies the list of resources and principals and their associated attributes that Verified Permissions can examine when evaluating the policies.")
	verifiedpermissions_batchIsAuthorizedCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store.")
	verifiedpermissions_batchIsAuthorizedCmd.Flags().String("requests", "", "An array of up to 30 requests that you want Verified Permissions to evaluate.")
	verifiedpermissions_batchIsAuthorizedCmd.MarkFlagRequired("policy-store-id")
	verifiedpermissions_batchIsAuthorizedCmd.MarkFlagRequired("requests")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_batchIsAuthorizedCmd)
}
