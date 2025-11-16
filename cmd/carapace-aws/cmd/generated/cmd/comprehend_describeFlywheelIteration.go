package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeFlywheelIterationCmd = &cobra.Command{
	Use:   "describe-flywheel-iteration",
	Short: "Retrieve the configuration properties of a flywheel iteration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeFlywheelIterationCmd).Standalone()

	comprehend_describeFlywheelIterationCmd.Flags().String("flywheel-arn", "", "")
	comprehend_describeFlywheelIterationCmd.Flags().String("flywheel-iteration-id", "", "")
	comprehend_describeFlywheelIterationCmd.MarkFlagRequired("flywheel-arn")
	comprehend_describeFlywheelIterationCmd.MarkFlagRequired("flywheel-iteration-id")
	comprehendCmd.AddCommand(comprehend_describeFlywheelIterationCmd)
}
