package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the given tags (metadata) from the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_untagResourceCmd).Standalone()

		iotfleetwise_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		iotfleetwise_untagResourceCmd.Flags().String("tag-keys", "", "A list of the keys of the tags to be removed from the resource.")
		iotfleetwise_untagResourceCmd.MarkFlagRequired("resource-arn")
		iotfleetwise_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_untagResourceCmd)
}
