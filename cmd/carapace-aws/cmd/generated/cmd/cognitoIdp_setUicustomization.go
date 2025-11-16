package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_setUicustomizationCmd = &cobra.Command{
	Use:   "set-uicustomization",
	Short: "Configures UI branding settings for domains with the hosted UI (classic) branding version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_setUicustomizationCmd).Standalone()

	cognitoIdp_setUicustomizationCmd.Flags().String("client-id", "", "The ID of the app client that you want to customize.")
	cognitoIdp_setUicustomizationCmd.Flags().String("css", "", "A plaintext CSS file that contains the custom fields that you want to apply to your user pool or app client.")
	cognitoIdp_setUicustomizationCmd.Flags().String("image-file", "", "The image that you want to set as your login in the classic hosted UI, as a Base64-formatted binary object.")
	cognitoIdp_setUicustomizationCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to apply branding to the classic hosted UI.")
	cognitoIdp_setUicustomizationCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_setUicustomizationCmd)
}
