package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_deleteProfilingGroupCmd = &cobra.Command{
	Use:   "delete-profiling-group",
	Short: "Deletes a profiling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_deleteProfilingGroupCmd).Standalone()

	codeguruprofiler_deleteProfilingGroupCmd.Flags().String("profiling-group-name", "", "The name of the profiling group to delete.")
	codeguruprofiler_deleteProfilingGroupCmd.MarkFlagRequired("profiling-group-name")
	codeguruprofilerCmd.AddCommand(codeguruprofiler_deleteProfilingGroupCmd)
}
