package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_setRulePrioritiesCmd = &cobra.Command{
	Use:   "set-rule-priorities",
	Short: "Sets the priorities of the specified rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_setRulePrioritiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_setRulePrioritiesCmd).Standalone()

		elbv2_setRulePrioritiesCmd.Flags().String("rule-priorities", "", "The rule priorities.")
		elbv2_setRulePrioritiesCmd.MarkFlagRequired("rule-priorities")
	})
	elbv2Cmd.AddCommand(elbv2_setRulePrioritiesCmd)
}
