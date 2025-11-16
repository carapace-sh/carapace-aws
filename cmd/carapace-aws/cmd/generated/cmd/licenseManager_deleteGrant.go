package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_deleteGrantCmd = &cobra.Command{
	Use:   "delete-grant",
	Short: "Deletes the specified grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_deleteGrantCmd).Standalone()

	licenseManager_deleteGrantCmd.Flags().String("grant-arn", "", "Amazon Resource Name (ARN) of the grant.")
	licenseManager_deleteGrantCmd.Flags().String("status-reason", "", "The Status reason for the delete request.")
	licenseManager_deleteGrantCmd.Flags().String("version", "", "Current version of the grant.")
	licenseManager_deleteGrantCmd.MarkFlagRequired("grant-arn")
	licenseManager_deleteGrantCmd.MarkFlagRequired("version")
	licenseManagerCmd.AddCommand(licenseManager_deleteGrantCmd)
}
