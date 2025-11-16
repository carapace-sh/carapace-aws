package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_listProfileTimesCmd = &cobra.Command{
	Use:   "list-profile-times",
	Short: "Lists the start times of the available aggregated profiles of a profiling group for an aggregation period within the specified time range.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_listProfileTimesCmd).Standalone()

	codeguruprofiler_listProfileTimesCmd.Flags().String("end-time", "", "The end time of the time range from which to list the profiles.")
	codeguruprofiler_listProfileTimesCmd.Flags().String("max-results", "", "The maximum number of profile time results returned by `ListProfileTimes` in paginated output.")
	codeguruprofiler_listProfileTimesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListProfileTimes` request where `maxResults` was used and the results exceeded the value of that parameter.")
	codeguruprofiler_listProfileTimesCmd.Flags().String("order-by", "", "The order (ascending or descending by start time of the profile) to use when listing profiles.")
	codeguruprofiler_listProfileTimesCmd.Flags().String("period", "", "The aggregation period.")
	codeguruprofiler_listProfileTimesCmd.Flags().String("profiling-group-name", "", "The name of the profiling group.")
	codeguruprofiler_listProfileTimesCmd.Flags().String("start-time", "", "The start time of the time range from which to list the profiles.")
	codeguruprofiler_listProfileTimesCmd.MarkFlagRequired("end-time")
	codeguruprofiler_listProfileTimesCmd.MarkFlagRequired("period")
	codeguruprofiler_listProfileTimesCmd.MarkFlagRequired("profiling-group-name")
	codeguruprofiler_listProfileTimesCmd.MarkFlagRequired("start-time")
	codeguruprofilerCmd.AddCommand(codeguruprofiler_listProfileTimesCmd)
}
