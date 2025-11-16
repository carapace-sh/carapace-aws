package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_untagResourceCmd).Standalone()

	iotwireless_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to remove tags from.")
	iotwireless_untagResourceCmd.Flags().String("tag-keys", "", "A list of the keys of the tags to remove from the resource.")
	iotwireless_untagResourceCmd.MarkFlagRequired("resource-arn")
	iotwireless_untagResourceCmd.MarkFlagRequired("tag-keys")
	iotwirelessCmd.AddCommand(iotwireless_untagResourceCmd)
}
