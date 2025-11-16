package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getBundlesCmd = &cobra.Command{
	Use:   "get-bundles",
	Short: "Returns the bundles that you can apply to an Amazon Lightsail instance when you create it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getBundlesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getBundlesCmd).Standalone()

		lightsail_getBundlesCmd.Flags().String("app-category", "", "Returns a list of bundles that are specific to Lightsail for Research.")
		lightsail_getBundlesCmd.Flags().String("include-inactive", "", "A Boolean value that indicates whether to include inactive (unavailable) bundles in the response of your request.")
		lightsail_getBundlesCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	})
	lightsailCmd.AddCommand(lightsail_getBundlesCmd)
}
