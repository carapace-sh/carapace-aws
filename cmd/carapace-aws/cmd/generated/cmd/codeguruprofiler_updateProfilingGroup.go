package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_updateProfilingGroupCmd = &cobra.Command{
	Use:   "update-profiling-group",
	Short: "Updates a profiling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_updateProfilingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruprofiler_updateProfilingGroupCmd).Standalone()

		codeguruprofiler_updateProfilingGroupCmd.Flags().String("agent-orchestration-config", "", "Specifies whether profiling is enabled or disabled for a profiling group.")
		codeguruprofiler_updateProfilingGroupCmd.Flags().String("profiling-group-name", "", "The name of the profiling group to update.")
		codeguruprofiler_updateProfilingGroupCmd.MarkFlagRequired("agent-orchestration-config")
		codeguruprofiler_updateProfilingGroupCmd.MarkFlagRequired("profiling-group-name")
	})
	codeguruprofilerCmd.AddCommand(codeguruprofiler_updateProfilingGroupCmd)
}
