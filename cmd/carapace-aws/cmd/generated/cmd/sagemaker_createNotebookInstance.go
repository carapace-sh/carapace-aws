package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createNotebookInstanceCmd = &cobra.Command{
	Use:   "create-notebook-instance",
	Short: "Creates an SageMaker AI notebook instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createNotebookInstanceCmd).Standalone()

	sagemaker_createNotebookInstanceCmd.Flags().String("accelerator-types", "", "This parameter is no longer supported.")
	sagemaker_createNotebookInstanceCmd.Flags().String("additional-code-repositories", "", "An array of up to three Git repositories to associate with the notebook instance.")
	sagemaker_createNotebookInstanceCmd.Flags().String("default-code-repository", "", "A Git repository to associate with the notebook instance as its default code repository.")
	sagemaker_createNotebookInstanceCmd.Flags().String("direct-internet-access", "", "Sets whether SageMaker AI provides internet access to the notebook instance.")
	sagemaker_createNotebookInstanceCmd.Flags().String("instance-metadata-service-configuration", "", "Information on the IMDS configuration of the notebook instance")
	sagemaker_createNotebookInstanceCmd.Flags().String("instance-type", "", "The type of ML compute instance to launch for the notebook instance.")
	sagemaker_createNotebookInstanceCmd.Flags().String("ip-address-type", "", "The IP address type for the notebook instance.")
	sagemaker_createNotebookInstanceCmd.Flags().String("kms-key-id", "", "The Amazon Resource Name (ARN) of a Amazon Web Services Key Management Service key that SageMaker AI uses to encrypt data on the storage volume attached to your notebook instance.")
	sagemaker_createNotebookInstanceCmd.Flags().String("lifecycle-config-name", "", "The name of a lifecycle configuration to associate with the notebook instance.")
	sagemaker_createNotebookInstanceCmd.Flags().String("notebook-instance-name", "", "The name of the new notebook instance.")
	sagemaker_createNotebookInstanceCmd.Flags().String("platform-identifier", "", "The platform identifier of the notebook instance runtime environment.")
	sagemaker_createNotebookInstanceCmd.Flags().String("role-arn", "", "When you send any requests to Amazon Web Services resources from the notebook instance, SageMaker AI assumes this role to perform tasks on your behalf.")
	sagemaker_createNotebookInstanceCmd.Flags().String("root-access", "", "Whether root access is enabled or disabled for users of the notebook instance.")
	sagemaker_createNotebookInstanceCmd.Flags().String("security-group-ids", "", "The VPC security group IDs, in the form sg-xxxxxxxx.")
	sagemaker_createNotebookInstanceCmd.Flags().String("subnet-id", "", "The ID of the subnet in a VPC to which you would like to have a connectivity from your ML compute instance.")
	sagemaker_createNotebookInstanceCmd.Flags().String("tags", "", "An array of key-value pairs.")
	sagemaker_createNotebookInstanceCmd.Flags().String("volume-size-in-gb", "", "The size, in GB, of the ML storage volume to attach to the notebook instance.")
	sagemaker_createNotebookInstanceCmd.MarkFlagRequired("instance-type")
	sagemaker_createNotebookInstanceCmd.MarkFlagRequired("notebook-instance-name")
	sagemaker_createNotebookInstanceCmd.MarkFlagRequired("role-arn")
	sagemakerCmd.AddCommand(sagemaker_createNotebookInstanceCmd)
}
