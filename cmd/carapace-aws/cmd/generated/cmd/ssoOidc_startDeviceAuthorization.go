package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoOidc_startDeviceAuthorizationCmd = &cobra.Command{
	Use:   "start-device-authorization",
	Short: "Initiates device authorization by requesting a pair of verification codes from the authorization service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoOidc_startDeviceAuthorizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoOidc_startDeviceAuthorizationCmd).Standalone()

		ssoOidc_startDeviceAuthorizationCmd.Flags().String("client-id", "", "The unique identifier string for the client that is registered with IAM Identity Center.")
		ssoOidc_startDeviceAuthorizationCmd.Flags().String("client-secret", "", "A secret string that is generated for the client.")
		ssoOidc_startDeviceAuthorizationCmd.Flags().String("start-url", "", "The URL for the Amazon Web Services access portal.")
		ssoOidc_startDeviceAuthorizationCmd.MarkFlagRequired("client-id")
		ssoOidc_startDeviceAuthorizationCmd.MarkFlagRequired("client-secret")
		ssoOidc_startDeviceAuthorizationCmd.MarkFlagRequired("start-url")
	})
	ssoOidcCmd.AddCommand(ssoOidc_startDeviceAuthorizationCmd)
}
