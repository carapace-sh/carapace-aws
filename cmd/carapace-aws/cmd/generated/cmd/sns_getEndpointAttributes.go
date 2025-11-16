package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_getEndpointAttributesCmd = &cobra.Command{
	Use:   "get-endpoint-attributes",
	Short: "Retrieves the endpoint attributes for a device on one of the supported push notification services, such as GCM (Firebase Cloud Messaging) and APNS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_getEndpointAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_getEndpointAttributesCmd).Standalone()

		sns_getEndpointAttributesCmd.Flags().String("endpoint-arn", "", "`EndpointArn` for `GetEndpointAttributes` input.")
		sns_getEndpointAttributesCmd.MarkFlagRequired("endpoint-arn")
	})
	snsCmd.AddCommand(sns_getEndpointAttributesCmd)
}
