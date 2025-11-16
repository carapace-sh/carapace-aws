package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_createProfilingGroupCmd = &cobra.Command{
	Use:   "create-profiling-group",
	Short: "Creates a profiling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_createProfilingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruprofiler_createProfilingGroupCmd).Standalone()

		codeguruprofiler_createProfilingGroupCmd.Flags().String("agent-orchestration-config", "", "Specifies whether profiling is enabled or disabled for the created profiling group.")
		codeguruprofiler_createProfilingGroupCmd.Flags().String("client-token", "", "Amazon CodeGuru Profiler uses this universally unique identifier (UUID) to prevent the accidental creation of duplicate profiling groups if there are failures and retries.")
		codeguruprofiler_createProfilingGroupCmd.Flags().String("compute-platform", "", "The compute platform of the profiling group.")
		codeguruprofiler_createProfilingGroupCmd.Flags().String("profiling-group-name", "", "The name of the profiling group to create.")
		codeguruprofiler_createProfilingGroupCmd.Flags().String("tags", "", "A list of tags to add to the created profiling group.")
		codeguruprofiler_createProfilingGroupCmd.MarkFlagRequired("client-token")
		codeguruprofiler_createProfilingGroupCmd.MarkFlagRequired("profiling-group-name")
	})
	codeguruprofilerCmd.AddCommand(codeguruprofiler_createProfilingGroupCmd)
}
