package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_addCustomAttributesCmd = &cobra.Command{
	Use:   "add-custom-attributes",
	Short: "Adds additional user attributes to the user pool schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_addCustomAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_addCustomAttributesCmd).Standalone()

		cognitoIdp_addCustomAttributesCmd.Flags().String("custom-attributes", "", "An array of custom attribute names and other properties.")
		cognitoIdp_addCustomAttributesCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to add custom attributes.")
		cognitoIdp_addCustomAttributesCmd.MarkFlagRequired("custom-attributes")
		cognitoIdp_addCustomAttributesCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_addCustomAttributesCmd)
}
