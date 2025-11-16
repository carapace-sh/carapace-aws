package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_startEventsDetectionJobCmd = &cobra.Command{
	Use:   "start-events-detection-job",
	Short: "Starts an asynchronous event detection job for a collection of documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_startEventsDetectionJobCmd).Standalone()

	comprehend_startEventsDetectionJobCmd.Flags().String("client-request-token", "", "An unique identifier for the request.")
	comprehend_startEventsDetectionJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.")
	comprehend_startEventsDetectionJobCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input data for the job.")
	comprehend_startEventsDetectionJobCmd.Flags().String("job-name", "", "The identifier of the events detection job.")
	comprehend_startEventsDetectionJobCmd.Flags().String("language-code", "", "The language code of the input documents.")
	comprehend_startEventsDetectionJobCmd.Flags().String("output-data-config", "", "Specifies where to send the output files.")
	comprehend_startEventsDetectionJobCmd.Flags().String("tags", "", "Tags to associate with the events detection job.")
	comprehend_startEventsDetectionJobCmd.Flags().String("target-event-types", "", "The types of events to detect in the input documents.")
	comprehend_startEventsDetectionJobCmd.MarkFlagRequired("data-access-role-arn")
	comprehend_startEventsDetectionJobCmd.MarkFlagRequired("input-data-config")
	comprehend_startEventsDetectionJobCmd.MarkFlagRequired("language-code")
	comprehend_startEventsDetectionJobCmd.MarkFlagRequired("output-data-config")
	comprehend_startEventsDetectionJobCmd.MarkFlagRequired("target-event-types")
	comprehendCmd.AddCommand(comprehend_startEventsDetectionJobCmd)
}
