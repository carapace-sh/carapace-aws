package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or overwrites the specified tags for the specified Amazon Managed Blockchain resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchain_tagResourceCmd).Standalone()

		managedblockchain_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		managedblockchain_tagResourceCmd.Flags().String("tags", "", "The tags to assign to the specified resource.")
		managedblockchain_tagResourceCmd.MarkFlagRequired("resource-arn")
		managedblockchain_tagResourceCmd.MarkFlagRequired("tags")
	})
	managedblockchainCmd.AddCommand(managedblockchain_tagResourceCmd)
}
