package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeFlywheelCmd = &cobra.Command{
	Use:   "describe-flywheel",
	Short: "Provides configuration information about the flywheel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeFlywheelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_describeFlywheelCmd).Standalone()

		comprehend_describeFlywheelCmd.Flags().String("flywheel-arn", "", "The Amazon Resource Number (ARN) of the flywheel.")
		comprehend_describeFlywheelCmd.MarkFlagRequired("flywheel-arn")
	})
	comprehendCmd.AddCommand(comprehend_describeFlywheelCmd)
}
