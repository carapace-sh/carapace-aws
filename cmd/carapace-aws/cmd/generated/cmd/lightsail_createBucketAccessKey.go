package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createBucketAccessKeyCmd = &cobra.Command{
	Use:   "create-bucket-access-key",
	Short: "Creates a new access key for the specified Amazon Lightsail bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createBucketAccessKeyCmd).Standalone()

	lightsail_createBucketAccessKeyCmd.Flags().String("bucket-name", "", "The name of the bucket that the new access key will belong to, and grant access to.")
	lightsail_createBucketAccessKeyCmd.MarkFlagRequired("bucket-name")
	lightsailCmd.AddCommand(lightsail_createBucketAccessKeyCmd)
}
