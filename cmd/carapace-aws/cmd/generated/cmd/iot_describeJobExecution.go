package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeJobExecutionCmd = &cobra.Command{
	Use:   "describe-job-execution",
	Short: "Describes a job execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeJobExecutionCmd).Standalone()

	iot_describeJobExecutionCmd.Flags().String("execution-number", "", "A string (consisting of the digits \"0\" through \"9\" which is used to specify a particular job execution on a particular device.")
	iot_describeJobExecutionCmd.Flags().String("job-id", "", "The unique identifier you assigned to this job when it was created.")
	iot_describeJobExecutionCmd.Flags().String("thing-name", "", "The name of the thing on which the job execution is running.")
	iot_describeJobExecutionCmd.MarkFlagRequired("job-id")
	iot_describeJobExecutionCmd.MarkFlagRequired("thing-name")
	iotCmd.AddCommand(iot_describeJobExecutionCmd)
}
