package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_getRegexPatternSetCmd = &cobra.Command{
	Use:   "get-regex-pattern-set",
	Short: "Retrieves the specified [RegexPatternSet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_getRegexPatternSetCmd).Standalone()

	wafv2_getRegexPatternSetCmd.Flags().String("id", "", "A unique identifier for the set.")
	wafv2_getRegexPatternSetCmd.Flags().String("name", "", "The name of the set.")
	wafv2_getRegexPatternSetCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_getRegexPatternSetCmd.MarkFlagRequired("id")
	wafv2_getRegexPatternSetCmd.MarkFlagRequired("name")
	wafv2_getRegexPatternSetCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_getRegexPatternSetCmd)
}
