package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_getGrantCmd = &cobra.Command{
	Use:   "get-grant",
	Short: "Gets detailed information about the specified grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_getGrantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_getGrantCmd).Standalone()

		licenseManager_getGrantCmd.Flags().String("grant-arn", "", "Amazon Resource Name (ARN) of the grant.")
		licenseManager_getGrantCmd.Flags().String("version", "", "Grant version.")
		licenseManager_getGrantCmd.MarkFlagRequired("grant-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_getGrantCmd)
}
