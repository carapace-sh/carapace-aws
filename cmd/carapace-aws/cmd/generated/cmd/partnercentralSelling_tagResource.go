package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_tagResourceCmd).Standalone()

	partnercentralSelling_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to tag.")
	partnercentralSelling_tagResourceCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign.")
	partnercentralSelling_tagResourceCmd.MarkFlagRequired("resource-arn")
	partnercentralSelling_tagResourceCmd.MarkFlagRequired("tags")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_tagResourceCmd)
}
