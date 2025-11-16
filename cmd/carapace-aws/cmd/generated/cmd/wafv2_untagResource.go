package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Disassociates tags from an Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_untagResourceCmd).Standalone()

		wafv2_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		wafv2_untagResourceCmd.Flags().String("tag-keys", "", "An array of keys identifying the tags to disassociate from the resource.")
		wafv2_untagResourceCmd.MarkFlagRequired("resource-arn")
		wafv2_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	wafv2Cmd.AddCommand(wafv2_untagResourceCmd)
}
