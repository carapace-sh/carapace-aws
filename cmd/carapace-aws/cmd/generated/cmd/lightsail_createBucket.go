package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createBucketCmd = &cobra.Command{
	Use:   "create-bucket",
	Short: "Creates an Amazon Lightsail bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createBucketCmd).Standalone()

	lightsail_createBucketCmd.Flags().String("bucket-name", "", "The name for the bucket.")
	lightsail_createBucketCmd.Flags().String("bundle-id", "", "The ID of the bundle to use for the bucket.")
	lightsail_createBucketCmd.Flags().String("enable-object-versioning", "", "A Boolean value that indicates whether to enable versioning of objects in the bucket.")
	lightsail_createBucketCmd.Flags().String("tags", "", "The tag keys and optional values to add to the bucket during creation.")
	lightsail_createBucketCmd.MarkFlagRequired("bucket-name")
	lightsail_createBucketCmd.MarkFlagRequired("bundle-id")
	lightsailCmd.AddCommand(lightsail_createBucketCmd)
}
