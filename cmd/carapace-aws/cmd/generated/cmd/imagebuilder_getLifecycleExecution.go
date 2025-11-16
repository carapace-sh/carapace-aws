package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getLifecycleExecutionCmd = &cobra.Command{
	Use:   "get-lifecycle-execution",
	Short: "Get the runtime information that was logged for a specific runtime instance of the lifecycle policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getLifecycleExecutionCmd).Standalone()

	imagebuilder_getLifecycleExecutionCmd.Flags().String("lifecycle-execution-id", "", "Use the unique identifier for a runtime instance of the lifecycle policy to get runtime details.")
	imagebuilder_getLifecycleExecutionCmd.MarkFlagRequired("lifecycle-execution-id")
	imagebuilderCmd.AddCommand(imagebuilder_getLifecycleExecutionCmd)
}
