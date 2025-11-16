package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_tagResourceCmd).Standalone()

		waf_tagResourceCmd.Flags().String("resource-arn", "", "")
		waf_tagResourceCmd.Flags().String("tags", "", "")
		waf_tagResourceCmd.MarkFlagRequired("resource-arn")
		waf_tagResourceCmd.MarkFlagRequired("tags")
	})
	wafCmd.AddCommand(waf_tagResourceCmd)
}
