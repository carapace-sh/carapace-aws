package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listJobExecutionsForJobCmd = &cobra.Command{
	Use:   "list-job-executions-for-job",
	Short: "Lists the job executions for a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listJobExecutionsForJobCmd).Standalone()

	iot_listJobExecutionsForJobCmd.Flags().String("job-id", "", "The unique identifier you assigned to this job when it was created.")
	iot_listJobExecutionsForJobCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	iot_listJobExecutionsForJobCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	iot_listJobExecutionsForJobCmd.Flags().String("status", "", "The status of the job.")
	iot_listJobExecutionsForJobCmd.MarkFlagRequired("job-id")
	iotCmd.AddCommand(iot_listJobExecutionsForJobCmd)
}
