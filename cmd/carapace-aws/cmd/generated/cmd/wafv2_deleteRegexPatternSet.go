package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_deleteRegexPatternSetCmd = &cobra.Command{
	Use:   "delete-regex-pattern-set",
	Short: "Deletes the specified [RegexPatternSet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_deleteRegexPatternSetCmd).Standalone()

	wafv2_deleteRegexPatternSetCmd.Flags().String("id", "", "A unique identifier for the set.")
	wafv2_deleteRegexPatternSetCmd.Flags().String("lock-token", "", "A token used for optimistic locking.")
	wafv2_deleteRegexPatternSetCmd.Flags().String("name", "", "The name of the set.")
	wafv2_deleteRegexPatternSetCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_deleteRegexPatternSetCmd.MarkFlagRequired("id")
	wafv2_deleteRegexPatternSetCmd.MarkFlagRequired("lock-token")
	wafv2_deleteRegexPatternSetCmd.MarkFlagRequired("name")
	wafv2_deleteRegexPatternSetCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_deleteRegexPatternSetCmd)
}
