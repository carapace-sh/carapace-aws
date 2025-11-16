package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_createPlatformEndpointCmd = &cobra.Command{
	Use:   "create-platform-endpoint",
	Short: "Creates an endpoint for a device and mobile app on one of the supported push notification services, such as GCM (Firebase Cloud Messaging) and APNS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_createPlatformEndpointCmd).Standalone()

	sns_createPlatformEndpointCmd.Flags().String("attributes", "", "For a list of attributes, see [`SetEndpointAttributes`](https://docs.aws.amazon.com/sns/latest/api/API_SetEndpointAttributes.html) .")
	sns_createPlatformEndpointCmd.Flags().String("custom-user-data", "", "Arbitrary user data to associate with the endpoint.")
	sns_createPlatformEndpointCmd.Flags().String("platform-application-arn", "", "`PlatformApplicationArn` returned from CreatePlatformApplication is used to create a an endpoint.")
	sns_createPlatformEndpointCmd.Flags().String("token", "", "Unique identifier created by the notification service for an app on a device.")
	sns_createPlatformEndpointCmd.MarkFlagRequired("platform-application-arn")
	sns_createPlatformEndpointCmd.MarkFlagRequired("token")
	snsCmd.AddCommand(sns_createPlatformEndpointCmd)
}
