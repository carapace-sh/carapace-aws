package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_isAuthorizedWithTokenCmd = &cobra.Command{
	Use:   "is-authorized-with-token",
	Short: "Makes an authorization decision about a service request described in the parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_isAuthorizedWithTokenCmd).Standalone()

	verifiedpermissions_isAuthorizedWithTokenCmd.Flags().String("access-token", "", "Specifies an access token for the principal to be authorized.")
	verifiedpermissions_isAuthorizedWithTokenCmd.Flags().String("action", "", "Specifies the requested action to be authorized.")
	verifiedpermissions_isAuthorizedWithTokenCmd.Flags().String("context", "", "Specifies additional context that can be used to make more granular authorization decisions.")
	verifiedpermissions_isAuthorizedWithTokenCmd.Flags().String("entities", "", "(Optional) Specifies the list of resources and their associated attributes that Verified Permissions can examine when evaluating the policies.")
	verifiedpermissions_isAuthorizedWithTokenCmd.Flags().String("identity-token", "", "Specifies an identity token for the principal to be authorized.")
	verifiedpermissions_isAuthorizedWithTokenCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store.")
	verifiedpermissions_isAuthorizedWithTokenCmd.Flags().String("resource", "", "Specifies the resource for which the authorization decision is made.")
	verifiedpermissions_isAuthorizedWithTokenCmd.MarkFlagRequired("policy-store-id")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_isAuthorizedWithTokenCmd)
}
