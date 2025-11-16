package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds to or modifies the tags of the given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_tagResourceCmd).Standalone()

		iotfleetwise_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		iotfleetwise_tagResourceCmd.Flags().String("tags", "", "The new or modified tags for the resource.")
		iotfleetwise_tagResourceCmd.MarkFlagRequired("resource-arn")
		iotfleetwise_tagResourceCmd.MarkFlagRequired("tags")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_tagResourceCmd)
}
