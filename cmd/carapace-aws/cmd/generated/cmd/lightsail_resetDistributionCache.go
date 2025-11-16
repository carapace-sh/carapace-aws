package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_resetDistributionCacheCmd = &cobra.Command{
	Use:   "reset-distribution-cache",
	Short: "Deletes currently cached content from your Amazon Lightsail content delivery network (CDN) distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_resetDistributionCacheCmd).Standalone()

	lightsail_resetDistributionCacheCmd.Flags().String("distribution-name", "", "The name of the distribution for which to reset cache.")
	lightsailCmd.AddCommand(lightsail_resetDistributionCacheCmd)
}
