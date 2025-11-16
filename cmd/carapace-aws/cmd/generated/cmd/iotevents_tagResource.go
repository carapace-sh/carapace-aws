package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds to or modifies the tags of the given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_tagResourceCmd).Standalone()

	iotevents_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	iotevents_tagResourceCmd.Flags().String("tags", "", "The new or modified tags for the resource.")
	iotevents_tagResourceCmd.MarkFlagRequired("resource-arn")
	iotevents_tagResourceCmd.MarkFlagRequired("tags")
	ioteventsCmd.AddCommand(iotevents_tagResourceCmd)
}
