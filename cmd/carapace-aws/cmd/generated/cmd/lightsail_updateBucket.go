package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_updateBucketCmd = &cobra.Command{
	Use:   "update-bucket",
	Short: "Updates an existing Amazon Lightsail bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_updateBucketCmd).Standalone()

	lightsail_updateBucketCmd.Flags().String("access-log-config", "", "An object that describes the access log configuration for the bucket.")
	lightsail_updateBucketCmd.Flags().String("access-rules", "", "An object that sets the public accessibility of objects in the specified bucket.")
	lightsail_updateBucketCmd.Flags().String("bucket-name", "", "The name of the bucket to update.")
	lightsail_updateBucketCmd.Flags().String("cors", "", "Sets the cross-origin resource sharing (CORS) configuration for your bucket.")
	lightsail_updateBucketCmd.Flags().String("readonly-access-accounts", "", "An array of strings to specify the Amazon Web Services account IDs that can access the bucket.")
	lightsail_updateBucketCmd.Flags().String("versioning", "", "Specifies whether to enable or suspend versioning of objects in the bucket.")
	lightsail_updateBucketCmd.MarkFlagRequired("bucket-name")
	lightsailCmd.AddCommand(lightsail_updateBucketCmd)
}
