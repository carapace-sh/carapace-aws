package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_listEndpointsByPlatformApplicationCmd = &cobra.Command{
	Use:   "list-endpoints-by-platform-application",
	Short: "Lists the endpoints and endpoint attributes for devices in a supported push notification service, such as GCM (Firebase Cloud Messaging) and APNS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_listEndpointsByPlatformApplicationCmd).Standalone()

	sns_listEndpointsByPlatformApplicationCmd.Flags().String("next-token", "", "`NextToken` string is used when calling `ListEndpointsByPlatformApplication` action to retrieve additional records that are available after the first page results.")
	sns_listEndpointsByPlatformApplicationCmd.Flags().String("platform-application-arn", "", "`PlatformApplicationArn` for `ListEndpointsByPlatformApplicationInput` action.")
	sns_listEndpointsByPlatformApplicationCmd.MarkFlagRequired("platform-application-arn")
	snsCmd.AddCommand(sns_listEndpointsByPlatformApplicationCmd)
}
