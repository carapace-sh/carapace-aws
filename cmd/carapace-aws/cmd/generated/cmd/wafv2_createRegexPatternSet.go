package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_createRegexPatternSetCmd = &cobra.Command{
	Use:   "create-regex-pattern-set",
	Short: "Creates a [RegexPatternSet](), which you reference in a [RegexPatternSetReferenceStatement](), to have WAF inspect a web request component for the specified patterns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_createRegexPatternSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_createRegexPatternSetCmd).Standalone()

		wafv2_createRegexPatternSetCmd.Flags().String("description", "", "A description of the set that helps with identification.")
		wafv2_createRegexPatternSetCmd.Flags().String("name", "", "The name of the set.")
		wafv2_createRegexPatternSetCmd.Flags().String("regular-expression-list", "", "Array of regular expression strings.")
		wafv2_createRegexPatternSetCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_createRegexPatternSetCmd.Flags().String("tags", "", "An array of key:value pairs to associate with the resource.")
		wafv2_createRegexPatternSetCmd.MarkFlagRequired("name")
		wafv2_createRegexPatternSetCmd.MarkFlagRequired("regular-expression-list")
		wafv2_createRegexPatternSetCmd.MarkFlagRequired("scope")
	})
	wafv2Cmd.AddCommand(wafv2_createRegexPatternSetCmd)
}
