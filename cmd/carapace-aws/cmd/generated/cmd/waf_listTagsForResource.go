package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_listTagsForResourceCmd).Standalone()

		waf_listTagsForResourceCmd.Flags().String("limit", "", "")
		waf_listTagsForResourceCmd.Flags().String("next-marker", "", "")
		waf_listTagsForResourceCmd.Flags().String("resource-arn", "", "")
		waf_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	wafCmd.AddCommand(waf_listTagsForResourceCmd)
}
