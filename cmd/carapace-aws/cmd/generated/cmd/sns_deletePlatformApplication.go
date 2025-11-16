package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_deletePlatformApplicationCmd = &cobra.Command{
	Use:   "delete-platform-application",
	Short: "Deletes a platform application object for one of the supported push notification services, such as APNS and GCM (Firebase Cloud Messaging).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_deletePlatformApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_deletePlatformApplicationCmd).Standalone()

		sns_deletePlatformApplicationCmd.Flags().String("platform-application-arn", "", "`PlatformApplicationArn` of platform application object to delete.")
		sns_deletePlatformApplicationCmd.MarkFlagRequired("platform-application-arn")
	})
	snsCmd.AddCommand(sns_deletePlatformApplicationCmd)
}
