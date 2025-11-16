package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_rejectGrantCmd = &cobra.Command{
	Use:   "reject-grant",
	Short: "Rejects the specified grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_rejectGrantCmd).Standalone()

	licenseManager_rejectGrantCmd.Flags().String("grant-arn", "", "Amazon Resource Name (ARN) of the grant.")
	licenseManager_rejectGrantCmd.MarkFlagRequired("grant-arn")
	licenseManagerCmd.AddCommand(licenseManager_rejectGrantCmd)
}
