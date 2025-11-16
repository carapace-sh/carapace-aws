package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_listRuleTypesCmd = &cobra.Command{
	Use:   "list-rule-types",
	Short: "Lists the rules for the condition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_listRuleTypesCmd).Standalone()

	codepipeline_listRuleTypesCmd.Flags().String("region-filter", "", "The rule Region to filter on.")
	codepipeline_listRuleTypesCmd.Flags().String("rule-owner-filter", "", "The rule owner to filter on.")
	codepipelineCmd.AddCommand(codepipeline_listRuleTypesCmd)
}
