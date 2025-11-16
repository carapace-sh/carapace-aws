package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_listMobileSdkReleasesCmd = &cobra.Command{
	Use:   "list-mobile-sdk-releases",
	Short: "Retrieves a list of the available releases for the mobile SDK and the specified device platform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_listMobileSdkReleasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_listMobileSdkReleasesCmd).Standalone()

		wafv2_listMobileSdkReleasesCmd.Flags().String("limit", "", "The maximum number of objects that you want WAF to return for this request.")
		wafv2_listMobileSdkReleasesCmd.Flags().String("next-marker", "", "When you request a list of objects with a `Limit` setting, if the number of objects that are still available for retrieval exceeds the limit, WAF returns a `NextMarker` value in the response.")
		wafv2_listMobileSdkReleasesCmd.Flags().String("platform", "", "The device platform to retrieve the list for.")
		wafv2_listMobileSdkReleasesCmd.MarkFlagRequired("platform")
	})
	wafv2Cmd.AddCommand(wafv2_listMobileSdkReleasesCmd)
}
