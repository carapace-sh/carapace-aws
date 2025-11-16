package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_updateIndexingRuleCmd = &cobra.Command{
	Use:   "update-indexing-rule",
	Short: "Modifies an indexing rule’s configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_updateIndexingRuleCmd).Standalone()

	xray_updateIndexingRuleCmd.Flags().String("name", "", "Name of the indexing rule to be updated.")
	xray_updateIndexingRuleCmd.Flags().String("rule", "", "Rule configuration to be updated.")
	xray_updateIndexingRuleCmd.MarkFlagRequired("name")
	xray_updateIndexingRuleCmd.MarkFlagRequired("rule")
	xrayCmd.AddCommand(xray_updateIndexingRuleCmd)
}
