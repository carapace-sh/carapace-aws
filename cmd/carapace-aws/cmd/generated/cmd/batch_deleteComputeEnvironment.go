package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_deleteComputeEnvironmentCmd = &cobra.Command{
	Use:   "delete-compute-environment",
	Short: "Deletes an Batch compute environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_deleteComputeEnvironmentCmd).Standalone()

	batch_deleteComputeEnvironmentCmd.Flags().String("compute-environment", "", "The name or Amazon Resource Name (ARN) of the compute environment to delete.")
	batch_deleteComputeEnvironmentCmd.MarkFlagRequired("compute-environment")
	batchCmd.AddCommand(batch_deleteComputeEnvironmentCmd)
}
