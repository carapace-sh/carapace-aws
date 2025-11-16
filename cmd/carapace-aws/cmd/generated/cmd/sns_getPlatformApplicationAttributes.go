package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_getPlatformApplicationAttributesCmd = &cobra.Command{
	Use:   "get-platform-application-attributes",
	Short: "Retrieves the attributes of the platform application object for the supported push notification services, such as APNS and GCM (Firebase Cloud Messaging).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_getPlatformApplicationAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_getPlatformApplicationAttributesCmd).Standalone()

		sns_getPlatformApplicationAttributesCmd.Flags().String("platform-application-arn", "", "`PlatformApplicationArn` for GetPlatformApplicationAttributesInput.")
		sns_getPlatformApplicationAttributesCmd.MarkFlagRequired("platform-application-arn")
	})
	snsCmd.AddCommand(sns_getPlatformApplicationAttributesCmd)
}
