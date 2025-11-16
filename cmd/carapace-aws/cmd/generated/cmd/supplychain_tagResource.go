package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "You can create tags during or after creating a resource such as instance, data flow, or dataset in AWS Supply chain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_tagResourceCmd).Standalone()

	supplychain_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Web Services Supply chain resource ARN that needs to be tagged.")
	supplychain_tagResourceCmd.Flags().String("tags", "", "The tags of the Amazon Web Services Supply chain resource to be created.")
	supplychain_tagResourceCmd.MarkFlagRequired("resource-arn")
	supplychain_tagResourceCmd.MarkFlagRequired("tags")
	supplychainCmd.AddCommand(supplychain_tagResourceCmd)
}
