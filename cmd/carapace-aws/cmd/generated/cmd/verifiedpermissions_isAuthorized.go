package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_isAuthorizedCmd = &cobra.Command{
	Use:   "is-authorized",
	Short: "Makes an authorization decision about a service request described in the parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_isAuthorizedCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_isAuthorizedCmd).Standalone()

		verifiedpermissions_isAuthorizedCmd.Flags().String("action", "", "Specifies the requested action to be authorized.")
		verifiedpermissions_isAuthorizedCmd.Flags().String("context", "", "Specifies additional context that can be used to make more granular authorization decisions.")
		verifiedpermissions_isAuthorizedCmd.Flags().String("entities", "", "(Optional) Specifies the list of resources and principals and their associated attributes that Verified Permissions can examine when evaluating the policies.")
		verifiedpermissions_isAuthorizedCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store.")
		verifiedpermissions_isAuthorizedCmd.Flags().String("principal", "", "Specifies the principal for which the authorization decision is to be made.")
		verifiedpermissions_isAuthorizedCmd.Flags().String("resource", "", "Specifies the resource for which the authorization decision is to be made.")
		verifiedpermissions_isAuthorizedCmd.MarkFlagRequired("policy-store-id")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_isAuthorizedCmd)
}
