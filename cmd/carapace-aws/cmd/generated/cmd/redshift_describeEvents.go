package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeEventsCmd = &cobra.Command{
	Use:   "describe-events",
	Short: "Returns events related to clusters, security groups, snapshots, and parameter groups for the past 14 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeEventsCmd).Standalone()

		redshift_describeEventsCmd.Flags().String("duration", "", "The number of minutes prior to the time of the request for which to retrieve events.")
		redshift_describeEventsCmd.Flags().String("end-time", "", "The end of the time interval for which to retrieve events, specified in ISO 8601 format.")
		redshift_describeEventsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
		redshift_describeEventsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
		redshift_describeEventsCmd.Flags().String("source-identifier", "", "The identifier of the event source for which events will be returned.")
		redshift_describeEventsCmd.Flags().String("source-type", "", "The event source to retrieve events for.")
		redshift_describeEventsCmd.Flags().String("start-time", "", "The beginning of the time interval to retrieve events for, specified in ISO 8601 format.")
	})
	redshiftCmd.AddCommand(redshift_describeEventsCmd)
}
