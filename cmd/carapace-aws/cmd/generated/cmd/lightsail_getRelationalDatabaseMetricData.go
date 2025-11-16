package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRelationalDatabaseMetricDataCmd = &cobra.Command{
	Use:   "get-relational-database-metric-data",
	Short: "Returns the data points of the specified metric for a database in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRelationalDatabaseMetricDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getRelationalDatabaseMetricDataCmd).Standalone()

		lightsail_getRelationalDatabaseMetricDataCmd.Flags().String("end-time", "", "The end of the time interval from which to get metric data.")
		lightsail_getRelationalDatabaseMetricDataCmd.Flags().String("metric-name", "", "The metric for which you want to return information.")
		lightsail_getRelationalDatabaseMetricDataCmd.Flags().String("period", "", "The granularity, in seconds, of the returned data points.")
		lightsail_getRelationalDatabaseMetricDataCmd.Flags().String("relational-database-name", "", "The name of your database from which to get metric data.")
		lightsail_getRelationalDatabaseMetricDataCmd.Flags().String("start-time", "", "The start of the time interval from which to get metric data.")
		lightsail_getRelationalDatabaseMetricDataCmd.Flags().String("statistics", "", "The statistic for the metric.")
		lightsail_getRelationalDatabaseMetricDataCmd.Flags().String("unit", "", "The unit for the metric data request.")
		lightsail_getRelationalDatabaseMetricDataCmd.MarkFlagRequired("end-time")
		lightsail_getRelationalDatabaseMetricDataCmd.MarkFlagRequired("metric-name")
		lightsail_getRelationalDatabaseMetricDataCmd.MarkFlagRequired("period")
		lightsail_getRelationalDatabaseMetricDataCmd.MarkFlagRequired("relational-database-name")
		lightsail_getRelationalDatabaseMetricDataCmd.MarkFlagRequired("start-time")
		lightsail_getRelationalDatabaseMetricDataCmd.MarkFlagRequired("statistics")
		lightsail_getRelationalDatabaseMetricDataCmd.MarkFlagRequired("unit")
	})
	lightsailCmd.AddCommand(lightsail_getRelationalDatabaseMetricDataCmd)
}
