package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateServicePipelineCmd = &cobra.Command{
	Use:   "update-service-pipeline",
	Short: "Update the service pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateServicePipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_updateServicePipelineCmd).Standalone()

		proton_updateServicePipelineCmd.Flags().String("deployment-type", "", "The deployment type.")
		proton_updateServicePipelineCmd.Flags().String("service-name", "", "The name of the service to that the pipeline is associated with.")
		proton_updateServicePipelineCmd.Flags().String("spec", "", "The spec for the service pipeline to update.")
		proton_updateServicePipelineCmd.Flags().String("template-major-version", "", "The major version of the service template that was used to create the service that the pipeline is associated with.")
		proton_updateServicePipelineCmd.Flags().String("template-minor-version", "", "The minor version of the service template that was used to create the service that the pipeline is associated with.")
		proton_updateServicePipelineCmd.MarkFlagRequired("deployment-type")
		proton_updateServicePipelineCmd.MarkFlagRequired("service-name")
		proton_updateServicePipelineCmd.MarkFlagRequired("spec")
	})
	protonCmd.AddCommand(proton_updateServicePipelineCmd)
}
