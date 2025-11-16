package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_startTestExecutionCmd = &cobra.Command{
	Use:   "start-test-execution",
	Short: "The action to start test set execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_startTestExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_startTestExecutionCmd).Standalone()

		lexv2Models_startTestExecutionCmd.Flags().String("api-mode", "", "Indicates whether we use streaming or non-streaming APIs for the test set execution.")
		lexv2Models_startTestExecutionCmd.Flags().String("target", "", "The target bot for the test set execution.")
		lexv2Models_startTestExecutionCmd.Flags().String("test-execution-modality", "", "Indicates whether audio or text is used.")
		lexv2Models_startTestExecutionCmd.Flags().String("test-set-id", "", "The test set Id for the test set execution.")
		lexv2Models_startTestExecutionCmd.MarkFlagRequired("api-mode")
		lexv2Models_startTestExecutionCmd.MarkFlagRequired("target")
		lexv2Models_startTestExecutionCmd.MarkFlagRequired("test-set-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_startTestExecutionCmd)
}
