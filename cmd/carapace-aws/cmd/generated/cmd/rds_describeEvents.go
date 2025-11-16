package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeEventsCmd = &cobra.Command{
	Use:   "describe-events",
	Short: "Returns events related to DB instances, DB clusters, DB parameter groups, DB security groups, DB snapshots, DB cluster snapshots, and RDS Proxies for the past 14 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeEventsCmd).Standalone()

		rds_describeEventsCmd.Flags().String("duration", "", "The number of minutes to retrieve events for.")
		rds_describeEventsCmd.Flags().String("end-time", "", "The end of the time interval for which to retrieve events, specified in ISO 8601 format.")
		rds_describeEventsCmd.Flags().String("event-categories", "", "A list of event categories that trigger notifications for a event notification subscription.")
		rds_describeEventsCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
		rds_describeEventsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous DescribeEvents request.")
		rds_describeEventsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeEventsCmd.Flags().String("source-identifier", "", "The identifier of the event source for which events are returned.")
		rds_describeEventsCmd.Flags().String("source-type", "", "The event source to retrieve events for.")
		rds_describeEventsCmd.Flags().String("start-time", "", "The beginning of the time interval to retrieve events for, specified in ISO 8601 format.")
	})
	rdsCmd.AddCommand(rds_describeEventsCmd)
}
