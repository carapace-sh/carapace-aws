package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_startFlywheelIterationCmd = &cobra.Command{
	Use:   "start-flywheel-iteration",
	Short: "Start the flywheel iteration.This operation uses any new datasets to train a new model version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_startFlywheelIterationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_startFlywheelIterationCmd).Standalone()

		comprehend_startFlywheelIterationCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
		comprehend_startFlywheelIterationCmd.Flags().String("flywheel-arn", "", "The ARN of the flywheel.")
		comprehend_startFlywheelIterationCmd.MarkFlagRequired("flywheel-arn")
	})
	comprehendCmd.AddCommand(comprehend_startFlywheelIterationCmd)
}
