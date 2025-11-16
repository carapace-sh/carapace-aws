package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_updateBucketBundleCmd = &cobra.Command{
	Use:   "update-bucket-bundle",
	Short: "Updates the bundle, or storage plan, of an existing Amazon Lightsail bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_updateBucketBundleCmd).Standalone()

	lightsail_updateBucketBundleCmd.Flags().String("bucket-name", "", "The name of the bucket for which to update the bundle.")
	lightsail_updateBucketBundleCmd.Flags().String("bundle-id", "", "The ID of the new bundle to apply to the bucket.")
	lightsail_updateBucketBundleCmd.MarkFlagRequired("bucket-name")
	lightsail_updateBucketBundleCmd.MarkFlagRequired("bundle-id")
	lightsailCmd.AddCommand(lightsail_updateBucketBundleCmd)
}
