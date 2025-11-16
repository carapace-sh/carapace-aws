package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_createTokenCmd = &cobra.Command{
	Use:   "create-token",
	Short: "Creates a long-lived token.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_createTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_createTokenCmd).Standalone()

		licenseManager_createTokenCmd.Flags().String("client-token", "", "Idempotency token, valid for 10 minutes.")
		licenseManager_createTokenCmd.Flags().String("expiration-in-days", "", "Token expiration, in days, counted from token creation.")
		licenseManager_createTokenCmd.Flags().String("license-arn", "", "Amazon Resource Name (ARN) of the license.")
		licenseManager_createTokenCmd.Flags().String("role-arns", "", "Amazon Resource Name (ARN) of the IAM roles to embed in the token.")
		licenseManager_createTokenCmd.Flags().String("token-properties", "", "Data specified by the caller to be included in the JWT token.")
		licenseManager_createTokenCmd.MarkFlagRequired("client-token")
		licenseManager_createTokenCmd.MarkFlagRequired("license-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_createTokenCmd)
}
