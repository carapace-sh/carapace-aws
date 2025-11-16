package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags (metadata) you have assigned to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_listTagsForResourceCmd).Standalone()

		iotwireless_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource for which you want to list tags.")
		iotwireless_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	iotwirelessCmd.AddCommand(iotwireless_listTagsForResourceCmd)
}
