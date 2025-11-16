package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_getUicustomizationCmd = &cobra.Command{
	Use:   "get-uicustomization",
	Short: "Given a user pool ID or app client, returns information about classic hosted UI branding that you applied, if any.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_getUicustomizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_getUicustomizationCmd).Standalone()

		cognitoIdp_getUicustomizationCmd.Flags().String("client-id", "", "The ID of the app client that you want to query for branding settings.")
		cognitoIdp_getUicustomizationCmd.Flags().String("user-pool-id", "", "The ID of the user pool that you want to query for branding settings.")
		cognitoIdp_getUicustomizationCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_getUicustomizationCmd)
}
