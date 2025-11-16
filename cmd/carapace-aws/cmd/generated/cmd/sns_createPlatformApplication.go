package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_createPlatformApplicationCmd = &cobra.Command{
	Use:   "create-platform-application",
	Short: "Creates a platform application object for one of the supported push notification services, such as APNS and GCM (Firebase Cloud Messaging), to which devices and mobile apps may register.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_createPlatformApplicationCmd).Standalone()

	sns_createPlatformApplicationCmd.Flags().String("attributes", "", "For a list of attributes, see [`SetPlatformApplicationAttributes`](https://docs.aws.amazon.com/sns/latest/api/API_SetPlatformApplicationAttributes.html) .")
	sns_createPlatformApplicationCmd.Flags().String("name", "", "Application names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, hyphens, and periods, and must be between 1 and 256 characters long.")
	sns_createPlatformApplicationCmd.Flags().String("platform", "", "The following platforms are supported: ADM (Amazon Device Messaging), APNS (Apple Push Notification Service), APNS\\_SANDBOX, and GCM (Firebase Cloud Messaging).")
	sns_createPlatformApplicationCmd.MarkFlagRequired("attributes")
	sns_createPlatformApplicationCmd.MarkFlagRequired("name")
	sns_createPlatformApplicationCmd.MarkFlagRequired("platform")
	snsCmd.AddCommand(sns_createPlatformApplicationCmd)
}
