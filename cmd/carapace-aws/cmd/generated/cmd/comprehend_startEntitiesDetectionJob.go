package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_startEntitiesDetectionJobCmd = &cobra.Command{
	Use:   "start-entities-detection-job",
	Short: "Starts an asynchronous entity detection job for a collection of documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_startEntitiesDetectionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_startEntitiesDetectionJobCmd).Standalone()

		comprehend_startEntitiesDetectionJobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
		comprehend_startEntitiesDetectionJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.")
		comprehend_startEntitiesDetectionJobCmd.Flags().String("entity-recognizer-arn", "", "The Amazon Resource Name (ARN) that identifies the specific entity recognizer to be used by the `StartEntitiesDetectionJob`.")
		comprehend_startEntitiesDetectionJobCmd.Flags().String("flywheel-arn", "", "The Amazon Resource Number (ARN) of the flywheel associated with the model to use.")
		comprehend_startEntitiesDetectionJobCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input data for the job.")
		comprehend_startEntitiesDetectionJobCmd.Flags().String("job-name", "", "The identifier of the job.")
		comprehend_startEntitiesDetectionJobCmd.Flags().String("language-code", "", "The language of the input documents.")
		comprehend_startEntitiesDetectionJobCmd.Flags().String("output-data-config", "", "Specifies where to send the output files.")
		comprehend_startEntitiesDetectionJobCmd.Flags().String("tags", "", "Tags to associate with the entities detection job.")
		comprehend_startEntitiesDetectionJobCmd.Flags().String("volume-kms-key-id", "", "ID for the Amazon Web Services Key Management Service (KMS) key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s) that process the analysis job.")
		comprehend_startEntitiesDetectionJobCmd.Flags().String("vpc-config", "", "Configuration parameters for an optional private Virtual Private Cloud (VPC) containing the resources you are using for your entity detection job.")
		comprehend_startEntitiesDetectionJobCmd.MarkFlagRequired("data-access-role-arn")
		comprehend_startEntitiesDetectionJobCmd.MarkFlagRequired("input-data-config")
		comprehend_startEntitiesDetectionJobCmd.MarkFlagRequired("language-code")
		comprehend_startEntitiesDetectionJobCmd.MarkFlagRequired("output-data-config")
	})
	comprehendCmd.AddCommand(comprehend_startEntitiesDetectionJobCmd)
}
