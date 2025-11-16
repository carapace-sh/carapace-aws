package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_updateRunGroupCmd = &cobra.Command{
	Use:   "update-run-group",
	Short: "Updates the settings of a run group and returns a response with no body if the operation is successful.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_updateRunGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_updateRunGroupCmd).Standalone()

		omics_updateRunGroupCmd.Flags().String("id", "", "The group's ID.")
		omics_updateRunGroupCmd.Flags().String("max-cpus", "", "The maximum number of CPUs to use.")
		omics_updateRunGroupCmd.Flags().String("max-duration", "", "A maximum run time for the group in minutes.")
		omics_updateRunGroupCmd.Flags().String("max-gpus", "", "The maximum GPUs that can be used by a run group.")
		omics_updateRunGroupCmd.Flags().String("max-runs", "", "The maximum number of concurrent runs for the group.")
		omics_updateRunGroupCmd.Flags().String("name", "", "A name for the group.")
		omics_updateRunGroupCmd.MarkFlagRequired("id")
	})
	omicsCmd.AddCommand(omics_updateRunGroupCmd)
}
