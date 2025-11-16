package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags (metadata) you have assigned to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotevents_listTagsForResourceCmd).Standalone()

		iotevents_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		iotevents_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	ioteventsCmd.AddCommand(iotevents_listTagsForResourceCmd)
}
