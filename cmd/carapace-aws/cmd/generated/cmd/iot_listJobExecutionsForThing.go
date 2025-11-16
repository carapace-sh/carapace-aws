package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listJobExecutionsForThingCmd = &cobra.Command{
	Use:   "list-job-executions-for-thing",
	Short: "Lists the job executions for the specified thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listJobExecutionsForThingCmd).Standalone()

	iot_listJobExecutionsForThingCmd.Flags().String("job-id", "", "The unique identifier you assigned to this job when it was created.")
	iot_listJobExecutionsForThingCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	iot_listJobExecutionsForThingCmd.Flags().String("namespace-id", "", "The namespace used to indicate that a job is a customer-managed job.")
	iot_listJobExecutionsForThingCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	iot_listJobExecutionsForThingCmd.Flags().String("status", "", "An optional filter that lets you search for jobs that have the specified status.")
	iot_listJobExecutionsForThingCmd.Flags().String("thing-name", "", "The thing name.")
	iot_listJobExecutionsForThingCmd.MarkFlagRequired("thing-name")
	iotCmd.AddCommand(iot_listJobExecutionsForThingCmd)
}
