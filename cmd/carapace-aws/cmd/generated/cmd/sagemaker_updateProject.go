package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateProjectCmd = &cobra.Command{
	Use:   "update-project",
	Short: "Updates a machine learning (ML) project that is created from a template that sets up an ML pipeline from training to deploying an approved model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateProjectCmd).Standalone()

	sagemaker_updateProjectCmd.Flags().String("project-description", "", "The description for the project.")
	sagemaker_updateProjectCmd.Flags().String("project-name", "", "The name of the project.")
	sagemaker_updateProjectCmd.Flags().String("service-catalog-provisioning-update-details", "", "The product ID and provisioning artifact ID to provision a service catalog.")
	sagemaker_updateProjectCmd.Flags().String("tags", "", "An array of key-value pairs.")
	sagemaker_updateProjectCmd.Flags().String("template-providers-to-update", "", "The template providers to update in the project.")
	sagemaker_updateProjectCmd.MarkFlagRequired("project-name")
	sagemakerCmd.AddCommand(sagemaker_updateProjectCmd)
}
