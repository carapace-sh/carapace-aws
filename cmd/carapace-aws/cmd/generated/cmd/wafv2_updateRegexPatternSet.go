package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_updateRegexPatternSetCmd = &cobra.Command{
	Use:   "update-regex-pattern-set",
	Short: "Updates the specified [RegexPatternSet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_updateRegexPatternSetCmd).Standalone()

	wafv2_updateRegexPatternSetCmd.Flags().String("description", "", "A description of the set that helps with identification.")
	wafv2_updateRegexPatternSetCmd.Flags().String("id", "", "A unique identifier for the set.")
	wafv2_updateRegexPatternSetCmd.Flags().String("lock-token", "", "A token used for optimistic locking.")
	wafv2_updateRegexPatternSetCmd.Flags().String("name", "", "The name of the set.")
	wafv2_updateRegexPatternSetCmd.Flags().String("regular-expression-list", "", "")
	wafv2_updateRegexPatternSetCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_updateRegexPatternSetCmd.MarkFlagRequired("id")
	wafv2_updateRegexPatternSetCmd.MarkFlagRequired("lock-token")
	wafv2_updateRegexPatternSetCmd.MarkFlagRequired("name")
	wafv2_updateRegexPatternSetCmd.MarkFlagRequired("regular-expression-list")
	wafv2_updateRegexPatternSetCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_updateRegexPatternSetCmd)
}
