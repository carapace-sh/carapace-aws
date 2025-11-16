package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes the specified set of tag keys and their values from the specified Amazon Lightsail resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_untagResourceCmd).Standalone()

		lightsail_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource from which you want to remove a tag.")
		lightsail_untagResourceCmd.Flags().String("resource-name", "", "The name of the resource from which you are removing a tag.")
		lightsail_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to delete from the specified resource.")
		lightsail_untagResourceCmd.MarkFlagRequired("resource-name")
		lightsail_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	lightsailCmd.AddCommand(lightsail_untagResourceCmd)
}
