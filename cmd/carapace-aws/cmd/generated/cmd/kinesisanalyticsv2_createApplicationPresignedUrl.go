package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_createApplicationPresignedUrlCmd = &cobra.Command{
	Use:   "create-application-presigned-url",
	Short: "Creates and returns a URL that you can use to connect to an application's extension.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_createApplicationPresignedUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_createApplicationPresignedUrlCmd).Standalone()

		kinesisanalyticsv2_createApplicationPresignedUrlCmd.Flags().String("application-name", "", "The name of the application.")
		kinesisanalyticsv2_createApplicationPresignedUrlCmd.Flags().String("session-expiration-duration-in-seconds", "", "The duration in seconds for which the returned URL will be valid.")
		kinesisanalyticsv2_createApplicationPresignedUrlCmd.Flags().String("url-type", "", "The type of the extension for which to create and return a URL.")
		kinesisanalyticsv2_createApplicationPresignedUrlCmd.MarkFlagRequired("application-name")
		kinesisanalyticsv2_createApplicationPresignedUrlCmd.MarkFlagRequired("url-type")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_createApplicationPresignedUrlCmd)
}
