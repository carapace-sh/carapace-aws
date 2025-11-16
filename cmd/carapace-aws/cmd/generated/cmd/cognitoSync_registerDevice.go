package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_registerDeviceCmd = &cobra.Command{
	Use:   "register-device",
	Short: "Registers a device to receive push sync notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_registerDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoSync_registerDeviceCmd).Standalone()

		cognitoSync_registerDeviceCmd.Flags().String("identity-id", "", "The unique ID for this identity.")
		cognitoSync_registerDeviceCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
		cognitoSync_registerDeviceCmd.Flags().String("platform", "", "The SNS platform type (e.g. GCM, SDM, APNS, APNS\\_SANDBOX).")
		cognitoSync_registerDeviceCmd.Flags().String("token", "", "The push token.")
		cognitoSync_registerDeviceCmd.MarkFlagRequired("identity-id")
		cognitoSync_registerDeviceCmd.MarkFlagRequired("identity-pool-id")
		cognitoSync_registerDeviceCmd.MarkFlagRequired("platform")
		cognitoSync_registerDeviceCmd.MarkFlagRequired("token")
	})
	cognitoSyncCmd.AddCommand(cognitoSync_registerDeviceCmd)
}
