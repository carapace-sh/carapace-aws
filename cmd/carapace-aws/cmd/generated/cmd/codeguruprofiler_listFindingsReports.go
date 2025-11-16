package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_listFindingsReportsCmd = &cobra.Command{
	Use:   "list-findings-reports",
	Short: "List the available reports for a given profiling group and time range.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_listFindingsReportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruprofiler_listFindingsReportsCmd).Standalone()

		codeguruprofiler_listFindingsReportsCmd.Flags().Bool("daily-reports-only", false, "A `Boolean` value indicating whether to only return reports from daily profiles.")
		codeguruprofiler_listFindingsReportsCmd.Flags().String("end-time", "", "The end time of the profile to get analysis data about.")
		codeguruprofiler_listFindingsReportsCmd.Flags().String("max-results", "", "The maximum number of report results returned by `ListFindingsReports` in paginated output.")
		codeguruprofiler_listFindingsReportsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListFindingsReportsRequest` request where `maxResults` was used and the results exceeded the value of that parameter.")
		codeguruprofiler_listFindingsReportsCmd.Flags().Bool("no-daily-reports-only", false, "A `Boolean` value indicating whether to only return reports from daily profiles.")
		codeguruprofiler_listFindingsReportsCmd.Flags().String("profiling-group-name", "", "The name of the profiling group from which to search for analysis data.")
		codeguruprofiler_listFindingsReportsCmd.Flags().String("start-time", "", "The start time of the profile to get analysis data about.")
		codeguruprofiler_listFindingsReportsCmd.MarkFlagRequired("end-time")
		codeguruprofiler_listFindingsReportsCmd.Flag("no-daily-reports-only").Hidden = true
		codeguruprofiler_listFindingsReportsCmd.MarkFlagRequired("profiling-group-name")
		codeguruprofiler_listFindingsReportsCmd.MarkFlagRequired("start-time")
	})
	codeguruprofilerCmd.AddCommand(codeguruprofiler_listFindingsReportsCmd)
}
