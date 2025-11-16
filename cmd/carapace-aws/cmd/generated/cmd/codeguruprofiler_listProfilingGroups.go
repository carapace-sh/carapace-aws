package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_listProfilingGroupsCmd = &cobra.Command{
	Use:   "list-profiling-groups",
	Short: "Returns a list of profiling groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_listProfilingGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruprofiler_listProfilingGroupsCmd).Standalone()

		codeguruprofiler_listProfilingGroupsCmd.Flags().Bool("include-description", false, "A `Boolean` value indicating whether to include a description.")
		codeguruprofiler_listProfilingGroupsCmd.Flags().String("max-results", "", "The maximum number of profiling groups results returned by `ListProfilingGroups` in paginated output.")
		codeguruprofiler_listProfilingGroupsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListProfilingGroups` request where `maxResults` was used and the results exceeded the value of that parameter.")
		codeguruprofiler_listProfilingGroupsCmd.Flags().Bool("no-include-description", false, "A `Boolean` value indicating whether to include a description.")
		codeguruprofiler_listProfilingGroupsCmd.Flag("no-include-description").Hidden = true
	})
	codeguruprofilerCmd.AddCommand(codeguruprofiler_listProfilingGroupsCmd)
}
