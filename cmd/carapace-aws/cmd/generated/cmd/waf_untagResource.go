package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_untagResourceCmd).Standalone()

	waf_untagResourceCmd.Flags().String("resource-arn", "", "")
	waf_untagResourceCmd.Flags().String("tag-keys", "", "")
	waf_untagResourceCmd.MarkFlagRequired("resource-arn")
	waf_untagResourceCmd.MarkFlagRequired("tag-keys")
	wafCmd.AddCommand(waf_untagResourceCmd)
}
