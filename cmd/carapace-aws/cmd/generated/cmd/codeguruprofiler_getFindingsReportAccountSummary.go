package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_getFindingsReportAccountSummaryCmd = &cobra.Command{
	Use:   "get-findings-report-account-summary",
	Short: "Returns a list of [`FindingsReportSummary`](https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_FindingsReportSummary.html) objects that contain analysis results for all profiling groups in your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_getFindingsReportAccountSummaryCmd).Standalone()

	codeguruprofiler_getFindingsReportAccountSummaryCmd.Flags().Bool("daily-reports-only", false, "A `Boolean` value indicating whether to only return reports from daily profiles.")
	codeguruprofiler_getFindingsReportAccountSummaryCmd.Flags().String("max-results", "", "The maximum number of results returned by `GetFindingsReportAccountSummary` in paginated output.")
	codeguruprofiler_getFindingsReportAccountSummaryCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `GetFindingsReportAccountSummary` request where `maxResults` was used and the results exceeded the value of that parameter.")
	codeguruprofiler_getFindingsReportAccountSummaryCmd.Flags().Bool("no-daily-reports-only", false, "A `Boolean` value indicating whether to only return reports from daily profiles.")
	codeguruprofiler_getFindingsReportAccountSummaryCmd.Flag("no-daily-reports-only").Hidden = true
	codeguruprofilerCmd.AddCommand(codeguruprofiler_getFindingsReportAccountSummaryCmd)
}
