package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_updateDistributionBundleCmd = &cobra.Command{
	Use:   "update-distribution-bundle",
	Short: "Updates the bundle of your Amazon Lightsail content delivery network (CDN) distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_updateDistributionBundleCmd).Standalone()

	lightsail_updateDistributionBundleCmd.Flags().String("bundle-id", "", "The bundle ID of the new bundle to apply to your distribution.")
	lightsail_updateDistributionBundleCmd.Flags().String("distribution-name", "", "The name of the distribution for which to update the bundle.")
	lightsailCmd.AddCommand(lightsail_updateDistributionBundleCmd)
}
