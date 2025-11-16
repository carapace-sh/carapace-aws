package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getBucketAccessKeysCmd = &cobra.Command{
	Use:   "get-bucket-access-keys",
	Short: "Returns the existing access key IDs for the specified Amazon Lightsail bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getBucketAccessKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getBucketAccessKeysCmd).Standalone()

		lightsail_getBucketAccessKeysCmd.Flags().String("bucket-name", "", "The name of the bucket for which to return access keys.")
		lightsail_getBucketAccessKeysCmd.MarkFlagRequired("bucket-name")
	})
	lightsailCmd.AddCommand(lightsail_getBucketAccessKeysCmd)
}
