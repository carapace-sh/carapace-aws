package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getDistributionLatestCacheResetCmd = &cobra.Command{
	Use:   "get-distribution-latest-cache-reset",
	Short: "Returns the timestamp and status of the last cache reset of a specific Amazon Lightsail content delivery network (CDN) distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getDistributionLatestCacheResetCmd).Standalone()

	lightsail_getDistributionLatestCacheResetCmd.Flags().String("distribution-name", "", "The name of the distribution for which to return the timestamp of the last cache reset.")
	lightsailCmd.AddCommand(lightsail_getDistributionLatestCacheResetCmd)
}
