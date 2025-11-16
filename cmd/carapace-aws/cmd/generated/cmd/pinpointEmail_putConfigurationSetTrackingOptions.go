package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_putConfigurationSetTrackingOptionsCmd = &cobra.Command{
	Use:   "put-configuration-set-tracking-options",
	Short: "Specify a custom domain to use for open and click tracking elements in email that you send using Amazon Pinpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_putConfigurationSetTrackingOptionsCmd).Standalone()

	pinpointEmail_putConfigurationSetTrackingOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that you want to add a custom tracking domain to.")
	pinpointEmail_putConfigurationSetTrackingOptionsCmd.Flags().String("custom-redirect-domain", "", "The domain that you want to use to track open and click events.")
	pinpointEmail_putConfigurationSetTrackingOptionsCmd.MarkFlagRequired("configuration-set-name")
	pinpointEmailCmd.AddCommand(pinpointEmail_putConfigurationSetTrackingOptionsCmd)
}
