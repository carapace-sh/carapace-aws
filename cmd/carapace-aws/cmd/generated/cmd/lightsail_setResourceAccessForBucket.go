package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_setResourceAccessForBucketCmd = &cobra.Command{
	Use:   "set-resource-access-for-bucket",
	Short: "Sets the Amazon Lightsail resources that can access the specified Lightsail bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_setResourceAccessForBucketCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_setResourceAccessForBucketCmd).Standalone()

		lightsail_setResourceAccessForBucketCmd.Flags().String("access", "", "The access setting.")
		lightsail_setResourceAccessForBucketCmd.Flags().String("bucket-name", "", "The name of the bucket for which to set access to another Lightsail resource.")
		lightsail_setResourceAccessForBucketCmd.Flags().String("resource-name", "", "The name of the Lightsail instance for which to set bucket access.")
		lightsail_setResourceAccessForBucketCmd.MarkFlagRequired("access")
		lightsail_setResourceAccessForBucketCmd.MarkFlagRequired("bucket-name")
		lightsail_setResourceAccessForBucketCmd.MarkFlagRequired("resource-name")
	})
	lightsailCmd.AddCommand(lightsail_setResourceAccessForBucketCmd)
}
