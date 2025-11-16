package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_listInsightsMetricDataCmd = &cobra.Command{
	Use:   "list-insights-metric-data",
	Short: "Returns Insights metrics data for trails that have enabled Insights.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_listInsightsMetricDataCmd).Standalone()

	cloudtrail_listInsightsMetricDataCmd.Flags().String("data-type", "", "Type of data points to return.")
	cloudtrail_listInsightsMetricDataCmd.Flags().String("end-time", "", "Specifies, in UTC, the end time for time-series data.")
	cloudtrail_listInsightsMetricDataCmd.Flags().String("error-code", "", "Conditionally required if the `InsightType` parameter is set to `ApiErrorRateInsight`.")
	cloudtrail_listInsightsMetricDataCmd.Flags().String("event-name", "", "The name of the event, typically the Amazon Web Services API on which unusual levels of activity were recorded.")
	cloudtrail_listInsightsMetricDataCmd.Flags().String("event-source", "", "The Amazon Web Services service to which the request was made, such as `iam.amazonaws.com` or `s3.amazonaws.com`.")
	cloudtrail_listInsightsMetricDataCmd.Flags().String("insight-type", "", "The type of CloudTrail Insights event, which is either `ApiCallRateInsight` or `ApiErrorRateInsight`.")
	cloudtrail_listInsightsMetricDataCmd.Flags().String("max-results", "", "The maximum number of data points to return.")
	cloudtrail_listInsightsMetricDataCmd.Flags().String("next-token", "", "Returned if all datapoints can't be returned in a single call.")
	cloudtrail_listInsightsMetricDataCmd.Flags().String("period", "", "Granularity of data to retrieve, in seconds.")
	cloudtrail_listInsightsMetricDataCmd.Flags().String("start-time", "", "Specifies, in UTC, the start time for time-series data.")
	cloudtrail_listInsightsMetricDataCmd.MarkFlagRequired("event-name")
	cloudtrail_listInsightsMetricDataCmd.MarkFlagRequired("event-source")
	cloudtrail_listInsightsMetricDataCmd.MarkFlagRequired("insight-type")
	cloudtrailCmd.AddCommand(cloudtrail_listInsightsMetricDataCmd)
}
