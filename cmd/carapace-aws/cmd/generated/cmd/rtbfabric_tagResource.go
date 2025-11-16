package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabric_tagResourceCmd).Standalone()

		rtbfabric_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to tag.")
		rtbfabric_tagResourceCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
		rtbfabric_tagResourceCmd.MarkFlagRequired("resource-arn")
		rtbfabric_tagResourceCmd.MarkFlagRequired("tags")
	})
	rtbfabricCmd.AddCommand(rtbfabric_tagResourceCmd)
}
