package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_startResourceEvaluationCmd = &cobra.Command{
	Use:   "start-resource-evaluation",
	Short: "Runs an on-demand evaluation for the specified resource to determine whether the resource details will comply with configured Config rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_startResourceEvaluationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_startResourceEvaluationCmd).Standalone()

		config_startResourceEvaluationCmd.Flags().String("client-token", "", "A client token is a unique, case-sensitive string of up to 64 ASCII characters.")
		config_startResourceEvaluationCmd.Flags().String("evaluation-context", "", "Returns an `EvaluationContext` object.")
		config_startResourceEvaluationCmd.Flags().String("evaluation-mode", "", "The mode of an evaluation.")
		config_startResourceEvaluationCmd.Flags().String("evaluation-timeout", "", "The timeout for an evaluation.")
		config_startResourceEvaluationCmd.Flags().String("resource-details", "", "Returns a `ResourceDetails` object.")
		config_startResourceEvaluationCmd.MarkFlagRequired("evaluation-mode")
		config_startResourceEvaluationCmd.MarkFlagRequired("resource-details")
	})
	configCmd.AddCommand(config_startResourceEvaluationCmd)
}
