package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getDistributionBundlesCmd = &cobra.Command{
	Use:   "get-distribution-bundles",
	Short: "Returns the bundles that can be applied to your Amazon Lightsail content delivery network (CDN) distributions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getDistributionBundlesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getDistributionBundlesCmd).Standalone()

	})
	lightsailCmd.AddCommand(lightsail_getDistributionBundlesCmd)
}
