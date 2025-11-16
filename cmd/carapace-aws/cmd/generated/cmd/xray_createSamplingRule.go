package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_createSamplingRuleCmd = &cobra.Command{
	Use:   "create-sampling-rule",
	Short: "Creates a rule to control sampling behavior for instrumented applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_createSamplingRuleCmd).Standalone()

	xray_createSamplingRuleCmd.Flags().String("sampling-rule", "", "The rule definition.")
	xray_createSamplingRuleCmd.Flags().String("tags", "", "A map that contains one or more tag keys and tag values to attach to an X-Ray sampling rule.")
	xray_createSamplingRuleCmd.MarkFlagRequired("sampling-rule")
	xrayCmd.AddCommand(xray_createSamplingRuleCmd)
}
