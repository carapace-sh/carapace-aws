package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Creates a machine learning (ML) project that can contain one or more templates that set up an ML pipeline from training to deploying an approved model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createProjectCmd).Standalone()

	sagemaker_createProjectCmd.Flags().String("project-description", "", "A description for the project.")
	sagemaker_createProjectCmd.Flags().String("project-name", "", "The name of the project.")
	sagemaker_createProjectCmd.Flags().String("service-catalog-provisioning-details", "", "The product ID and provisioning artifact ID to provision a service catalog.")
	sagemaker_createProjectCmd.Flags().String("tags", "", "An array of key-value pairs that you want to use to organize and track your Amazon Web Services resource costs.")
	sagemaker_createProjectCmd.Flags().String("template-providers", "", "An array of template provider configurations for creating infrastructure resources for the project.")
	sagemaker_createProjectCmd.MarkFlagRequired("project-name")
	sagemakerCmd.AddCommand(sagemaker_createProjectCmd)
}
