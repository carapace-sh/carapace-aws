package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_listTagsForResourceCmd).Standalone()

	managedblockchain_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	managedblockchain_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	managedblockchainCmd.AddCommand(managedblockchain_listTagsForResourceCmd)
}
