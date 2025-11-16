package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_getPolicyStoreCmd = &cobra.Command{
	Use:   "get-policy-store",
	Short: "Retrieves details about a policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_getPolicyStoreCmd).Standalone()

	verifiedpermissions_getPolicyStoreCmd.Flags().Bool("no-tags", false, "Specifies whether to return the tags that are attached to the policy store.")
	verifiedpermissions_getPolicyStoreCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that you want information about.")
	verifiedpermissions_getPolicyStoreCmd.Flags().Bool("tags", false, "Specifies whether to return the tags that are attached to the policy store.")
	verifiedpermissions_getPolicyStoreCmd.Flag("no-tags").Hidden = true
	verifiedpermissions_getPolicyStoreCmd.MarkFlagRequired("policy-store-id")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_getPolicyStoreCmd)
}
