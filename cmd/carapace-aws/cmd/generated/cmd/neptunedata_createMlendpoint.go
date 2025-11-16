package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_createMlendpointCmd = &cobra.Command{
	Use:   "create-mlendpoint",
	Short: "Creates a new Neptune ML inference endpoint that lets you query one specific model that the model-training process constructed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_createMlendpointCmd).Standalone()

	neptunedata_createMlendpointCmd.Flags().String("id", "", "A unique identifier for the new inference endpoint.")
	neptunedata_createMlendpointCmd.Flags().String("instance-count", "", "The minimum number of Amazon EC2 instances to deploy to an endpoint for prediction.")
	neptunedata_createMlendpointCmd.Flags().String("instance-type", "", "The type of Neptune ML instance to use for online servicing.")
	neptunedata_createMlendpointCmd.Flags().String("ml-model-training-job-id", "", "The job Id of the completed model-training job that has created the model that the inference endpoint will point to.")
	neptunedata_createMlendpointCmd.Flags().String("ml-model-transform-job-id", "", "The job Id of the completed model-transform job.")
	neptunedata_createMlendpointCmd.Flags().String("model-name", "", "Model type for training.")
	neptunedata_createMlendpointCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role providing Neptune access to SageMaker and Amazon S3 resources.")
	neptunedata_createMlendpointCmd.Flags().Bool("no-update", false, "If set to `true`, `update` indicates that this is an update request.")
	neptunedata_createMlendpointCmd.Flags().Bool("update", false, "If set to `true`, `update` indicates that this is an update request.")
	neptunedata_createMlendpointCmd.Flags().String("volume-encryption-kmskey", "", "The Amazon Key Management Service (Amazon KMS) key that SageMaker uses to encrypt data on the storage volume attached to the ML compute instances that run the training job.")
	neptunedata_createMlendpointCmd.Flag("no-update").Hidden = true
	neptunedataCmd.AddCommand(neptunedata_createMlendpointCmd)
}
