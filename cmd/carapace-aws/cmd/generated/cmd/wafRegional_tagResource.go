package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_tagResourceCmd).Standalone()

		wafRegional_tagResourceCmd.Flags().String("resource-arn", "", "")
		wafRegional_tagResourceCmd.Flags().String("tags", "", "")
		wafRegional_tagResourceCmd.MarkFlagRequired("resource-arn")
		wafRegional_tagResourceCmd.MarkFlagRequired("tags")
	})
	wafRegionalCmd.AddCommand(wafRegional_tagResourceCmd)
}
