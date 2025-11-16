package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_acceptGrantCmd = &cobra.Command{
	Use:   "accept-grant",
	Short: "Accepts the specified grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_acceptGrantCmd).Standalone()

	licenseManager_acceptGrantCmd.Flags().String("grant-arn", "", "Amazon Resource Name (ARN) of the grant.")
	licenseManager_acceptGrantCmd.MarkFlagRequired("grant-arn")
	licenseManagerCmd.AddCommand(licenseManager_acceptGrantCmd)
}
