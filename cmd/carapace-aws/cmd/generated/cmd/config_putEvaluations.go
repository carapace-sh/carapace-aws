package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putEvaluationsCmd = &cobra.Command{
	Use:   "put-evaluations",
	Short: "Used by an Lambda function to deliver evaluation results to Config.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putEvaluationsCmd).Standalone()

	config_putEvaluationsCmd.Flags().String("evaluations", "", "The assessments that the Lambda function performs.")
	config_putEvaluationsCmd.Flags().Bool("no-test-mode", false, "Use this parameter to specify a test run for `PutEvaluations`.")
	config_putEvaluationsCmd.Flags().String("result-token", "", "An encrypted token that associates an evaluation with an Config rule.")
	config_putEvaluationsCmd.Flags().Bool("test-mode", false, "Use this parameter to specify a test run for `PutEvaluations`.")
	config_putEvaluationsCmd.Flag("no-test-mode").Hidden = true
	config_putEvaluationsCmd.MarkFlagRequired("result-token")
	configCmd.AddCommand(config_putEvaluationsCmd)
}
