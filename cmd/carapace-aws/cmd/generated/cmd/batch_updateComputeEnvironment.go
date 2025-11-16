package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_updateComputeEnvironmentCmd = &cobra.Command{
	Use:   "update-compute-environment",
	Short: "Updates an Batch compute environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_updateComputeEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_updateComputeEnvironmentCmd).Standalone()

		batch_updateComputeEnvironmentCmd.Flags().String("compute-environment", "", "The name or full Amazon Resource Name (ARN) of the compute environment to update.")
		batch_updateComputeEnvironmentCmd.Flags().String("compute-resources", "", "Details of the compute resources managed by the compute environment.")
		batch_updateComputeEnvironmentCmd.Flags().String("context", "", "Reserved.")
		batch_updateComputeEnvironmentCmd.Flags().String("service-role", "", "The full Amazon Resource Name (ARN) of the IAM role that allows Batch to make calls to other Amazon Web Services services on your behalf.")
		batch_updateComputeEnvironmentCmd.Flags().String("state", "", "The state of the compute environment.")
		batch_updateComputeEnvironmentCmd.Flags().String("unmanagedv-cpus", "", "The maximum number of vCPUs expected to be used for an unmanaged compute environment.")
		batch_updateComputeEnvironmentCmd.Flags().String("update-policy", "", "Specifies the updated infrastructure update policy for the compute environment.")
		batch_updateComputeEnvironmentCmd.MarkFlagRequired("compute-environment")
	})
	batchCmd.AddCommand(batch_updateComputeEnvironmentCmd)
}
