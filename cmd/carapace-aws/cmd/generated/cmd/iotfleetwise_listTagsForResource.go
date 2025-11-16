package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags (metadata) you have assigned to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listTagsForResourceCmd).Standalone()

	iotfleetwise_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	iotfleetwise_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	iotfleetwiseCmd.AddCommand(iotfleetwise_listTagsForResourceCmd)
}
