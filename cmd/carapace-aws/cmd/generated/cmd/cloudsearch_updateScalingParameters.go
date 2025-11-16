package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_updateScalingParametersCmd = &cobra.Command{
	Use:   "update-scaling-parameters",
	Short: "Configures scaling parameters for a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_updateScalingParametersCmd).Standalone()

	cloudsearch_updateScalingParametersCmd.Flags().String("domain-name", "", "")
	cloudsearch_updateScalingParametersCmd.Flags().String("scaling-parameters", "", "")
	cloudsearch_updateScalingParametersCmd.MarkFlagRequired("domain-name")
	cloudsearch_updateScalingParametersCmd.MarkFlagRequired("scaling-parameters")
	cloudsearchCmd.AddCommand(cloudsearch_updateScalingParametersCmd)
}
