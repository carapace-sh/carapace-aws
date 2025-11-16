package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_getTestExecutionArtifactsUrlCmd = &cobra.Command{
	Use:   "get-test-execution-artifacts-url",
	Short: "The pre-signed Amazon S3 URL to download the test execution result artifacts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_getTestExecutionArtifactsUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_getTestExecutionArtifactsUrlCmd).Standalone()

		lexv2Models_getTestExecutionArtifactsUrlCmd.Flags().String("test-execution-id", "", "The unique identifier of the completed test execution.")
		lexv2Models_getTestExecutionArtifactsUrlCmd.MarkFlagRequired("test-execution-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_getTestExecutionArtifactsUrlCmd)
}
