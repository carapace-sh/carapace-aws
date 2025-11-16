package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_describeProfilingGroupCmd = &cobra.Command{
	Use:   "describe-profiling-group",
	Short: "Returns a [`ProfilingGroupDescription`](https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ProfilingGroupDescription.html) object that contains information about the requested profiling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_describeProfilingGroupCmd).Standalone()

	codeguruprofiler_describeProfilingGroupCmd.Flags().String("profiling-group-name", "", "The name of the profiling group to get information about.")
	codeguruprofiler_describeProfilingGroupCmd.MarkFlagRequired("profiling-group-name")
	codeguruprofilerCmd.AddCommand(codeguruprofiler_describeProfilingGroupCmd)
}
