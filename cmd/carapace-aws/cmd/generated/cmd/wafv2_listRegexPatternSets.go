package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_listRegexPatternSetsCmd = &cobra.Command{
	Use:   "list-regex-pattern-sets",
	Short: "Retrieves an array of [RegexPatternSetSummary]() objects for the regex pattern sets that you manage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_listRegexPatternSetsCmd).Standalone()

	wafv2_listRegexPatternSetsCmd.Flags().String("limit", "", "The maximum number of objects that you want WAF to return for this request.")
	wafv2_listRegexPatternSetsCmd.Flags().String("next-marker", "", "When you request a list of objects with a `Limit` setting, if the number of objects that are still available for retrieval exceeds the limit, WAF returns a `NextMarker` value in the response.")
	wafv2_listRegexPatternSetsCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_listRegexPatternSetsCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_listRegexPatternSetsCmd)
}
