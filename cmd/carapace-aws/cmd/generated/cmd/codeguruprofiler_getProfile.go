package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_getProfileCmd = &cobra.Command{
	Use:   "get-profile",
	Short: "Gets the aggregated profile of a profiling group for a specified time range.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_getProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruprofiler_getProfileCmd).Standalone()

		codeguruprofiler_getProfileCmd.Flags().String("accept", "", "The format of the returned profiling data.")
		codeguruprofiler_getProfileCmd.Flags().String("end-time", "", "The end time of the requested profile.")
		codeguruprofiler_getProfileCmd.Flags().String("max-depth", "", "The maximum depth of the stacks in the code that is represented in the aggregated profile.")
		codeguruprofiler_getProfileCmd.Flags().String("period", "", "Used with `startTime` or `endTime` to specify the time range for the returned aggregated profile.")
		codeguruprofiler_getProfileCmd.Flags().String("profiling-group-name", "", "The name of the profiling group to get.")
		codeguruprofiler_getProfileCmd.Flags().String("start-time", "", "The start time of the profile to get.")
		codeguruprofiler_getProfileCmd.MarkFlagRequired("profiling-group-name")
	})
	codeguruprofilerCmd.AddCommand(codeguruprofiler_getProfileCmd)
}
