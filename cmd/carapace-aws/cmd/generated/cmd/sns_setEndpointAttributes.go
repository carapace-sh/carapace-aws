package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_setEndpointAttributesCmd = &cobra.Command{
	Use:   "set-endpoint-attributes",
	Short: "Sets the attributes for an endpoint for a device on one of the supported push notification services, such as GCM (Firebase Cloud Messaging) and APNS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_setEndpointAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_setEndpointAttributesCmd).Standalone()

		sns_setEndpointAttributesCmd.Flags().String("attributes", "", "A map of the endpoint attributes.")
		sns_setEndpointAttributesCmd.Flags().String("endpoint-arn", "", "EndpointArn used for `SetEndpointAttributes` action.")
		sns_setEndpointAttributesCmd.MarkFlagRequired("attributes")
		sns_setEndpointAttributesCmd.MarkFlagRequired("endpoint-arn")
	})
	snsCmd.AddCommand(sns_setEndpointAttributesCmd)
}
