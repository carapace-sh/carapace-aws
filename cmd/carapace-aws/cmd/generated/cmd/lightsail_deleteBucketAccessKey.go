package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteBucketAccessKeyCmd = &cobra.Command{
	Use:   "delete-bucket-access-key",
	Short: "Deletes an access key for the specified Amazon Lightsail bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteBucketAccessKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_deleteBucketAccessKeyCmd).Standalone()

		lightsail_deleteBucketAccessKeyCmd.Flags().String("access-key-id", "", "The ID of the access key to delete.")
		lightsail_deleteBucketAccessKeyCmd.Flags().String("bucket-name", "", "The name of the bucket that the access key belongs to.")
		lightsail_deleteBucketAccessKeyCmd.MarkFlagRequired("access-key-id")
		lightsail_deleteBucketAccessKeyCmd.MarkFlagRequired("bucket-name")
	})
	lightsailCmd.AddCommand(lightsail_deleteBucketAccessKeyCmd)
}
