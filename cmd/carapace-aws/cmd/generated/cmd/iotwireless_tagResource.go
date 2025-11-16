package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a tag to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_tagResourceCmd).Standalone()

	iotwireless_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to add tags to.")
	iotwireless_tagResourceCmd.Flags().String("tags", "", "Adds to or modifies the tags of the given resource.")
	iotwireless_tagResourceCmd.MarkFlagRequired("resource-arn")
	iotwireless_tagResourceCmd.MarkFlagRequired("tags")
	iotwirelessCmd.AddCommand(iotwireless_tagResourceCmd)
}
