package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getBucketBundlesCmd = &cobra.Command{
	Use:   "get-bucket-bundles",
	Short: "Returns the bundles that you can apply to a Amazon Lightsail bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getBucketBundlesCmd).Standalone()

	lightsail_getBucketBundlesCmd.Flags().String("include-inactive", "", "A Boolean value that indicates whether to include inactive (unavailable) bundles in the response of your request.")
	lightsailCmd.AddCommand(lightsail_getBucketBundlesCmd)
}
