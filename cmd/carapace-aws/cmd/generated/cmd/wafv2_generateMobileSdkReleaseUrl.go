package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_generateMobileSdkReleaseUrlCmd = &cobra.Command{
	Use:   "generate-mobile-sdk-release-url",
	Short: "Generates a presigned download URL for the specified release of the mobile SDK.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_generateMobileSdkReleaseUrlCmd).Standalone()

	wafv2_generateMobileSdkReleaseUrlCmd.Flags().String("platform", "", "The device platform.")
	wafv2_generateMobileSdkReleaseUrlCmd.Flags().String("release-version", "", "The release version.")
	wafv2_generateMobileSdkReleaseUrlCmd.MarkFlagRequired("platform")
	wafv2_generateMobileSdkReleaseUrlCmd.MarkFlagRequired("release-version")
	wafv2Cmd.AddCommand(wafv2_generateMobileSdkReleaseUrlCmd)
}
