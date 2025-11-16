package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateNotebookInstanceCmd = &cobra.Command{
	Use:   "update-notebook-instance",
	Short: "Updates a notebook instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateNotebookInstanceCmd).Standalone()

	sagemaker_updateNotebookInstanceCmd.Flags().String("accelerator-types", "", "This parameter is no longer supported.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("additional-code-repositories", "", "An array of up to three Git repositories to associate with the notebook instance.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("default-code-repository", "", "The Git repository to associate with the notebook instance as its default code repository.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("disassociate-accelerator-types", "", "This parameter is no longer supported.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("disassociate-additional-code-repositories", "", "A list of names or URLs of the default Git repositories to remove from this notebook instance.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("disassociate-default-code-repository", "", "The name or URL of the default Git repository to remove from this notebook instance.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("disassociate-lifecycle-config", "", "Set to `true` to remove the notebook instance lifecycle configuration currently associated with the notebook instance.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("instance-metadata-service-configuration", "", "Information on the IMDS configuration of the notebook instance")
	sagemaker_updateNotebookInstanceCmd.Flags().String("instance-type", "", "The Amazon ML compute instance type.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("ip-address-type", "", "The IP address type for the notebook instance.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("lifecycle-config-name", "", "The name of a lifecycle configuration to associate with the notebook instance.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("notebook-instance-name", "", "The name of the notebook instance to update.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("platform-identifier", "", "The platform identifier of the notebook instance runtime environment.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that SageMaker AI can assume to access the notebook instance.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("root-access", "", "Whether root access is enabled or disabled for users of the notebook instance.")
	sagemaker_updateNotebookInstanceCmd.Flags().String("volume-size-in-gb", "", "The size, in GB, of the ML storage volume to attach to the notebook instance.")
	sagemaker_updateNotebookInstanceCmd.MarkFlagRequired("notebook-instance-name")
	sagemakerCmd.AddCommand(sagemaker_updateNotebookInstanceCmd)
}
