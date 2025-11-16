package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_listManagedRuleSetsCmd = &cobra.Command{
	Use:   "list-managed-rule-sets",
	Short: "Retrieves the managed rule sets that you own.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_listManagedRuleSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_listManagedRuleSetsCmd).Standalone()

		wafv2_listManagedRuleSetsCmd.Flags().String("limit", "", "The maximum number of objects that you want WAF to return for this request.")
		wafv2_listManagedRuleSetsCmd.Flags().String("next-marker", "", "When you request a list of objects with a `Limit` setting, if the number of objects that are still available for retrieval exceeds the limit, WAF returns a `NextMarker` value in the response.")
		wafv2_listManagedRuleSetsCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_listManagedRuleSetsCmd.MarkFlagRequired("scope")
	})
	wafv2Cmd.AddCommand(wafv2_listManagedRuleSetsCmd)
}
