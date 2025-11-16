package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_untagResourceCmd).Standalone()

		wafRegional_untagResourceCmd.Flags().String("resource-arn", "", "")
		wafRegional_untagResourceCmd.Flags().String("tag-keys", "", "")
		wafRegional_untagResourceCmd.MarkFlagRequired("resource-arn")
		wafRegional_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	wafRegionalCmd.AddCommand(wafRegional_untagResourceCmd)
}
