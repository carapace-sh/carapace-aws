package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "You can delete tags for an Amazon Web Services Supply chain resource such as instance, data flow, or dataset in AWS Supply Chain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_untagResourceCmd).Standalone()

		supplychain_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Web Services Supply chain resource ARN that needs to be untagged.")
		supplychain_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to be deleted for an Amazon Web Services Supply Chain resource.")
		supplychain_untagResourceCmd.MarkFlagRequired("resource-arn")
		supplychain_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	supplychainCmd.AddCommand(supplychain_untagResourceCmd)
}
