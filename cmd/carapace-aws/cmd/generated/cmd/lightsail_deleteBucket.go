package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteBucketCmd = &cobra.Command{
	Use:   "delete-bucket",
	Short: "Deletes a Amazon Lightsail bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteBucketCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_deleteBucketCmd).Standalone()

		lightsail_deleteBucketCmd.Flags().String("bucket-name", "", "The name of the bucket to delete.")
		lightsail_deleteBucketCmd.Flags().String("force-delete", "", "A Boolean value that indicates whether to force delete the bucket.")
		lightsail_deleteBucketCmd.MarkFlagRequired("bucket-name")
	})
	lightsailCmd.AddCommand(lightsail_deleteBucketCmd)
}
