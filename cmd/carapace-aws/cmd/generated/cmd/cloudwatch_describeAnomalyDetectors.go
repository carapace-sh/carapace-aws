package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_describeAnomalyDetectorsCmd = &cobra.Command{
	Use:   "describe-anomaly-detectors",
	Short: "Lists the anomaly detection models that you have created in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_describeAnomalyDetectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_describeAnomalyDetectorsCmd).Standalone()

		cloudwatch_describeAnomalyDetectorsCmd.Flags().String("anomaly-detector-types", "", "The anomaly detector types to request when using `DescribeAnomalyDetectorsInput`.")
		cloudwatch_describeAnomalyDetectorsCmd.Flags().String("dimensions", "", "Limits the results to only the anomaly detection models that are associated with the specified metric dimensions.")
		cloudwatch_describeAnomalyDetectorsCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
		cloudwatch_describeAnomalyDetectorsCmd.Flags().String("metric-name", "", "Limits the results to only the anomaly detection models that are associated with the specified metric name.")
		cloudwatch_describeAnomalyDetectorsCmd.Flags().String("namespace", "", "Limits the results to only the anomaly detection models that are associated with the specified namespace.")
		cloudwatch_describeAnomalyDetectorsCmd.Flags().String("next-token", "", "Use the token returned by the previous operation to request the next page of results.")
	})
	cloudwatchCmd.AddCommand(cloudwatch_describeAnomalyDetectorsCmd)
}
