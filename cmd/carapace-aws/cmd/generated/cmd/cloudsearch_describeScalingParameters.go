package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_describeScalingParametersCmd = &cobra.Command{
	Use:   "describe-scaling-parameters",
	Short: "Gets the scaling parameters configured for a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_describeScalingParametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_describeScalingParametersCmd).Standalone()

		cloudsearch_describeScalingParametersCmd.Flags().String("domain-name", "", "")
		cloudsearch_describeScalingParametersCmd.MarkFlagRequired("domain-name")
	})
	cloudsearchCmd.AddCommand(cloudsearch_describeScalingParametersCmd)
}
