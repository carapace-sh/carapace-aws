package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_getMobileSdkReleaseCmd = &cobra.Command{
	Use:   "get-mobile-sdk-release",
	Short: "Retrieves information for the specified mobile SDK release, including release notes and tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_getMobileSdkReleaseCmd).Standalone()

	wafv2_getMobileSdkReleaseCmd.Flags().String("platform", "", "The device platform.")
	wafv2_getMobileSdkReleaseCmd.Flags().String("release-version", "", "The release version.")
	wafv2_getMobileSdkReleaseCmd.MarkFlagRequired("platform")
	wafv2_getMobileSdkReleaseCmd.MarkFlagRequired("release-version")
	wafv2Cmd.AddCommand(wafv2_getMobileSdkReleaseCmd)
}
