package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_startTopicsDetectionJobCmd = &cobra.Command{
	Use:   "start-topics-detection-job",
	Short: "Starts an asynchronous topic detection job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_startTopicsDetectionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_startTopicsDetectionJobCmd).Standalone()

		comprehend_startTopicsDetectionJobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
		comprehend_startTopicsDetectionJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.")
		comprehend_startTopicsDetectionJobCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input data for the job.")
		comprehend_startTopicsDetectionJobCmd.Flags().String("job-name", "", "The identifier of the job.")
		comprehend_startTopicsDetectionJobCmd.Flags().String("number-of-topics", "", "The number of topics to detect.")
		comprehend_startTopicsDetectionJobCmd.Flags().String("output-data-config", "", "Specifies where to send the output files.")
		comprehend_startTopicsDetectionJobCmd.Flags().String("tags", "", "Tags to associate with the topics detection job.")
		comprehend_startTopicsDetectionJobCmd.Flags().String("volume-kms-key-id", "", "ID for the Amazon Web Services Key Management Service (KMS) key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s) that process the analysis job.")
		comprehend_startTopicsDetectionJobCmd.Flags().String("vpc-config", "", "Configuration parameters for an optional private Virtual Private Cloud (VPC) containing the resources you are using for your topic detection job.")
		comprehend_startTopicsDetectionJobCmd.MarkFlagRequired("data-access-role-arn")
		comprehend_startTopicsDetectionJobCmd.MarkFlagRequired("input-data-config")
		comprehend_startTopicsDetectionJobCmd.MarkFlagRequired("output-data-config")
	})
	comprehendCmd.AddCommand(comprehend_startTopicsDetectionJobCmd)
}
