package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all the tags for a specified resource, including key and value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_listTagsForResourceCmd).Standalone()

		lookoutequipment_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource (such as the dataset or model) that is the focus of the `ListTagsForResource` operation.")
		lookoutequipment_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_listTagsForResourceCmd)
}
