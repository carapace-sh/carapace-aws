package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove resource tags from a Greengrass Resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_untagResourceCmd).Standalone()

	greengrass_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	greengrass_untagResourceCmd.Flags().String("tag-keys", "", "An array of tag keys to delete")
	greengrass_untagResourceCmd.MarkFlagRequired("resource-arn")
	greengrass_untagResourceCmd.MarkFlagRequired("tag-keys")
	greengrassCmd.AddCommand(greengrass_untagResourceCmd)
}
