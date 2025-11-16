package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List all the tags for an Amazon Web ServicesSupply Chain resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_listTagsForResourceCmd).Standalone()

	supplychain_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Web Services Supply chain resource ARN that needs tags to be listed.")
	supplychain_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	supplychainCmd.AddCommand(supplychain_listTagsForResourceCmd)
}
