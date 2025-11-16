package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_createGrantCmd = &cobra.Command{
	Use:   "create-grant",
	Short: "Creates a grant for the specified license.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_createGrantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_createGrantCmd).Standalone()

		licenseManager_createGrantCmd.Flags().String("allowed-operations", "", "Allowed operations for the grant.")
		licenseManager_createGrantCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		licenseManager_createGrantCmd.Flags().String("grant-name", "", "Grant name.")
		licenseManager_createGrantCmd.Flags().String("home-region", "", "Home Region of the grant.")
		licenseManager_createGrantCmd.Flags().String("license-arn", "", "Amazon Resource Name (ARN) of the license.")
		licenseManager_createGrantCmd.Flags().String("principals", "", "The grant principals.")
		licenseManager_createGrantCmd.Flags().String("tags", "", "Tags to add to the grant.")
		licenseManager_createGrantCmd.MarkFlagRequired("allowed-operations")
		licenseManager_createGrantCmd.MarkFlagRequired("client-token")
		licenseManager_createGrantCmd.MarkFlagRequired("grant-name")
		licenseManager_createGrantCmd.MarkFlagRequired("home-region")
		licenseManager_createGrantCmd.MarkFlagRequired("license-arn")
		licenseManager_createGrantCmd.MarkFlagRequired("principals")
	})
	licenseManagerCmd.AddCommand(licenseManager_createGrantCmd)
}
