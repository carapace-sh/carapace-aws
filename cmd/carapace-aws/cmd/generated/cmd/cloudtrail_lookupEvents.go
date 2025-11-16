package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_lookupEventsCmd = &cobra.Command{
	Use:   "lookup-events",
	Short: "Looks up [management events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-management-events) or [CloudTrail Insights events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-insights-events) that are captured by CloudTrail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_lookupEventsCmd).Standalone()

	cloudtrail_lookupEventsCmd.Flags().String("end-time", "", "Specifies that only events that occur before or at the specified time are returned.")
	cloudtrail_lookupEventsCmd.Flags().String("event-category", "", "Specifies the event category.")
	cloudtrail_lookupEventsCmd.Flags().String("lookup-attributes", "", "Contains a list of lookup attributes.")
	cloudtrail_lookupEventsCmd.Flags().String("max-results", "", "The number of events to return.")
	cloudtrail_lookupEventsCmd.Flags().String("next-token", "", "The token to use to get the next page of results after a previous API call.")
	cloudtrail_lookupEventsCmd.Flags().String("start-time", "", "Specifies that only events that occur after or at the specified time are returned.")
	cloudtrailCmd.AddCommand(cloudtrail_lookupEventsCmd)
}
