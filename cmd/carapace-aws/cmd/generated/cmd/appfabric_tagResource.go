package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_tagResourceCmd).Standalone()

	appfabric_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to tag.")
	appfabric_tagResourceCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
	appfabric_tagResourceCmd.MarkFlagRequired("resource-arn")
	appfabric_tagResourceCmd.MarkFlagRequired("tags")
	appfabricCmd.AddCommand(appfabric_tagResourceCmd)
}
