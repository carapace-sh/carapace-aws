package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listTagsForResourceCmd).Standalone()

	wafRegional_listTagsForResourceCmd.Flags().String("limit", "", "")
	wafRegional_listTagsForResourceCmd.Flags().String("next-marker", "", "")
	wafRegional_listTagsForResourceCmd.Flags().String("resource-arn", "", "")
	wafRegional_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	wafRegionalCmd.AddCommand(wafRegional_listTagsForResourceCmd)
}
