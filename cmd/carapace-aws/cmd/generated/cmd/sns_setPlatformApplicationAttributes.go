package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_setPlatformApplicationAttributesCmd = &cobra.Command{
	Use:   "set-platform-application-attributes",
	Short: "Sets the attributes of the platform application object for the supported push notification services, such as APNS and GCM (Firebase Cloud Messaging).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_setPlatformApplicationAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_setPlatformApplicationAttributesCmd).Standalone()

		sns_setPlatformApplicationAttributesCmd.Flags().String("attributes", "", "A map of the platform application attributes.")
		sns_setPlatformApplicationAttributesCmd.Flags().String("platform-application-arn", "", "`PlatformApplicationArn` for `SetPlatformApplicationAttributes` action.")
		sns_setPlatformApplicationAttributesCmd.MarkFlagRequired("attributes")
		sns_setPlatformApplicationAttributesCmd.MarkFlagRequired("platform-application-arn")
	})
	snsCmd.AddCommand(sns_setPlatformApplicationAttributesCmd)
}
