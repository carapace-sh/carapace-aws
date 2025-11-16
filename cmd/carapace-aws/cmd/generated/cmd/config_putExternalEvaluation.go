package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putExternalEvaluationCmd = &cobra.Command{
	Use:   "put-external-evaluation",
	Short: "Add or updates the evaluations for process checks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putExternalEvaluationCmd).Standalone()

	config_putExternalEvaluationCmd.Flags().String("config-rule-name", "", "The name of the Config rule.")
	config_putExternalEvaluationCmd.Flags().String("external-evaluation", "", "An `ExternalEvaluation` object that provides details about compliance.")
	config_putExternalEvaluationCmd.MarkFlagRequired("config-rule-name")
	config_putExternalEvaluationCmd.MarkFlagRequired("external-evaluation")
	configCmd.AddCommand(config_putExternalEvaluationCmd)
}
