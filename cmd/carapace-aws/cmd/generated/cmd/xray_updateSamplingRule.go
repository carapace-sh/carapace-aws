package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_updateSamplingRuleCmd = &cobra.Command{
	Use:   "update-sampling-rule",
	Short: "Modifies a sampling rule's configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_updateSamplingRuleCmd).Standalone()

	xray_updateSamplingRuleCmd.Flags().String("sampling-rule-update", "", "The rule and fields to change.")
	xray_updateSamplingRuleCmd.MarkFlagRequired("sampling-rule-update")
	xrayCmd.AddCommand(xray_updateSamplingRuleCmd)
}
