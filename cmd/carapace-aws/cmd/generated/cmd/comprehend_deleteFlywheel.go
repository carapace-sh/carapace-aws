package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_deleteFlywheelCmd = &cobra.Command{
	Use:   "delete-flywheel",
	Short: "Deletes a flywheel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_deleteFlywheelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_deleteFlywheelCmd).Standalone()

		comprehend_deleteFlywheelCmd.Flags().String("flywheel-arn", "", "The Amazon Resource Number (ARN) of the flywheel to delete.")
		comprehend_deleteFlywheelCmd.MarkFlagRequired("flywheel-arn")
	})
	comprehendCmd.AddCommand(comprehend_deleteFlywheelCmd)
}
