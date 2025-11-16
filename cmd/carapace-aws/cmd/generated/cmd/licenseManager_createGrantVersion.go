package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_createGrantVersionCmd = &cobra.Command{
	Use:   "create-grant-version",
	Short: "Creates a new version of the specified grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_createGrantVersionCmd).Standalone()

	licenseManager_createGrantVersionCmd.Flags().String("allowed-operations", "", "Allowed operations for the grant.")
	licenseManager_createGrantVersionCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	licenseManager_createGrantVersionCmd.Flags().String("grant-arn", "", "Amazon Resource Name (ARN) of the grant.")
	licenseManager_createGrantVersionCmd.Flags().String("grant-name", "", "Grant name.")
	licenseManager_createGrantVersionCmd.Flags().String("options", "", "The options specified for the grant.")
	licenseManager_createGrantVersionCmd.Flags().String("source-version", "", "Current version of the grant.")
	licenseManager_createGrantVersionCmd.Flags().String("status", "", "Grant status.")
	licenseManager_createGrantVersionCmd.Flags().String("status-reason", "", "Grant status reason.")
	licenseManager_createGrantVersionCmd.MarkFlagRequired("client-token")
	licenseManager_createGrantVersionCmd.MarkFlagRequired("grant-arn")
	licenseManagerCmd.AddCommand(licenseManager_createGrantVersionCmd)
}
