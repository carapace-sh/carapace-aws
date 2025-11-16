package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_getAccessTokenCmd = &cobra.Command{
	Use:   "get-access-token",
	Short: "Gets a temporary access token to use with AssumeRoleWithWebIdentity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_getAccessTokenCmd).Standalone()

	licenseManager_getAccessTokenCmd.Flags().String("token", "", "Refresh token, encoded as a JWT token.")
	licenseManager_getAccessTokenCmd.Flags().String("token-properties", "", "Token properties to validate against those present in the JWT token.")
	licenseManager_getAccessTokenCmd.MarkFlagRequired("token")
	licenseManagerCmd.AddCommand(licenseManager_getAccessTokenCmd)
}
