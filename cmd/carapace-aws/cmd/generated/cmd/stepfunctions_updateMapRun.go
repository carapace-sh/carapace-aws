package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_updateMapRunCmd = &cobra.Command{
	Use:   "update-map-run",
	Short: "Updates an in-progress Map Run's configuration to include changes to the settings that control maximum concurrency and Map Run failure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_updateMapRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_updateMapRunCmd).Standalone()

		stepfunctions_updateMapRunCmd.Flags().String("map-run-arn", "", "The Amazon Resource Name (ARN) of a Map Run.")
		stepfunctions_updateMapRunCmd.Flags().String("max-concurrency", "", "The maximum number of child workflow executions that can be specified to run in parallel for the Map Run at the same time.")
		stepfunctions_updateMapRunCmd.Flags().String("tolerated-failure-count", "", "The maximum number of failed items before the Map Run fails.")
		stepfunctions_updateMapRunCmd.Flags().String("tolerated-failure-percentage", "", "The maximum percentage of failed items before the Map Run fails.")
		stepfunctions_updateMapRunCmd.MarkFlagRequired("map-run-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_updateMapRunCmd)
}
