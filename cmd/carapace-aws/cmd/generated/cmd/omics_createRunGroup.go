package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_createRunGroupCmd = &cobra.Command{
	Use:   "create-run-group",
	Short: "Creates a run group to limit the compute resources for the runs that are added to the group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_createRunGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_createRunGroupCmd).Standalone()

		omics_createRunGroupCmd.Flags().String("max-cpus", "", "The maximum number of CPUs that can run concurrently across all active runs in the run group.")
		omics_createRunGroupCmd.Flags().String("max-duration", "", "The maximum time for each run (in minutes).")
		omics_createRunGroupCmd.Flags().String("max-gpus", "", "The maximum number of GPUs that can run concurrently across all active runs in the run group.")
		omics_createRunGroupCmd.Flags().String("max-runs", "", "The maximum number of runs that can be running at the same time.")
		omics_createRunGroupCmd.Flags().String("name", "", "A name for the group.")
		omics_createRunGroupCmd.Flags().String("request-id", "", "To ensure that requests don't run multiple times, specify a unique ID for each request.")
		omics_createRunGroupCmd.Flags().String("tags", "", "Tags for the group.")
		omics_createRunGroupCmd.MarkFlagRequired("request-id")
	})
	omicsCmd.AddCommand(omics_createRunGroupCmd)
}
