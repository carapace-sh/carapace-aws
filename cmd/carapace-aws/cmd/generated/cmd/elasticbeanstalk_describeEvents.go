package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describeEventsCmd = &cobra.Command{
	Use:   "describe-events",
	Short: "Returns list of event descriptions matching criteria up to the last 6 weeks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describeEventsCmd).Standalone()

	elasticbeanstalk_describeEventsCmd.Flags().String("application-name", "", "If specified, AWS Elastic Beanstalk restricts the returned descriptions to include only those associated with this application.")
	elasticbeanstalk_describeEventsCmd.Flags().String("end-time", "", "If specified, AWS Elastic Beanstalk restricts the returned descriptions to those that occur up to, but not including, the `EndTime`.")
	elasticbeanstalk_describeEventsCmd.Flags().String("environment-id", "", "If specified, AWS Elastic Beanstalk restricts the returned descriptions to those associated with this environment.")
	elasticbeanstalk_describeEventsCmd.Flags().String("environment-name", "", "If specified, AWS Elastic Beanstalk restricts the returned descriptions to those associated with this environment.")
	elasticbeanstalk_describeEventsCmd.Flags().String("max-records", "", "Specifies the maximum number of events that can be returned, beginning with the most recent event.")
	elasticbeanstalk_describeEventsCmd.Flags().String("next-token", "", "Pagination token.")
	elasticbeanstalk_describeEventsCmd.Flags().String("platform-arn", "", "The ARN of a custom platform version.")
	elasticbeanstalk_describeEventsCmd.Flags().String("request-id", "", "If specified, AWS Elastic Beanstalk restricts the described events to include only those associated with this request ID.")
	elasticbeanstalk_describeEventsCmd.Flags().String("severity", "", "If specified, limits the events returned from this call to include only those with the specified severity or higher.")
	elasticbeanstalk_describeEventsCmd.Flags().String("start-time", "", "If specified, AWS Elastic Beanstalk restricts the returned descriptions to those that occur on or after this time.")
	elasticbeanstalk_describeEventsCmd.Flags().String("template-name", "", "If specified, AWS Elastic Beanstalk restricts the returned descriptions to those that are associated with this environment configuration.")
	elasticbeanstalk_describeEventsCmd.Flags().String("version-label", "", "If specified, AWS Elastic Beanstalk restricts the returned descriptions to those associated with this application version.")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describeEventsCmd)
}
