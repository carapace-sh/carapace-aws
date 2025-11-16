package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_startPiiEntitiesDetectionJobCmd = &cobra.Command{
	Use:   "start-pii-entities-detection-job",
	Short: "Starts an asynchronous PII entity detection job for a collection of documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_startPiiEntitiesDetectionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_startPiiEntitiesDetectionJobCmd).Standalone()

		comprehend_startPiiEntitiesDetectionJobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
		comprehend_startPiiEntitiesDetectionJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.")
		comprehend_startPiiEntitiesDetectionJobCmd.Flags().String("input-data-config", "", "The input properties for a PII entities detection job.")
		comprehend_startPiiEntitiesDetectionJobCmd.Flags().String("job-name", "", "The identifier of the job.")
		comprehend_startPiiEntitiesDetectionJobCmd.Flags().String("language-code", "", "The language of the input documents.")
		comprehend_startPiiEntitiesDetectionJobCmd.Flags().String("mode", "", "Specifies whether the output provides the locations (offsets) of PII entities or a file in which PII entities are redacted.")
		comprehend_startPiiEntitiesDetectionJobCmd.Flags().String("output-data-config", "", "Provides conﬁguration parameters for the output of PII entity detection jobs.")
		comprehend_startPiiEntitiesDetectionJobCmd.Flags().String("redaction-config", "", "Provides configuration parameters for PII entity redaction.")
		comprehend_startPiiEntitiesDetectionJobCmd.Flags().String("tags", "", "Tags to associate with the PII entities detection job.")
		comprehend_startPiiEntitiesDetectionJobCmd.MarkFlagRequired("data-access-role-arn")
		comprehend_startPiiEntitiesDetectionJobCmd.MarkFlagRequired("input-data-config")
		comprehend_startPiiEntitiesDetectionJobCmd.MarkFlagRequired("language-code")
		comprehend_startPiiEntitiesDetectionJobCmd.MarkFlagRequired("mode")
		comprehend_startPiiEntitiesDetectionJobCmd.MarkFlagRequired("output-data-config")
	})
	comprehendCmd.AddCommand(comprehend_startPiiEntitiesDetectionJobCmd)
}
