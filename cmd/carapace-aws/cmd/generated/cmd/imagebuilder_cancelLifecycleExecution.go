package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_cancelLifecycleExecutionCmd = &cobra.Command{
	Use:   "cancel-lifecycle-execution",
	Short: "Cancel a specific image lifecycle policy runtime instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_cancelLifecycleExecutionCmd).Standalone()

	imagebuilder_cancelLifecycleExecutionCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
	imagebuilder_cancelLifecycleExecutionCmd.Flags().String("lifecycle-execution-id", "", "Identifies the specific runtime instance of the image lifecycle to cancel.")
	imagebuilder_cancelLifecycleExecutionCmd.MarkFlagRequired("client-token")
	imagebuilder_cancelLifecycleExecutionCmd.MarkFlagRequired("lifecycle-execution-id")
	imagebuilderCmd.AddCommand(imagebuilder_cancelLifecycleExecutionCmd)
}
