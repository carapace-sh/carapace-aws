package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listServicePipelineOutputsCmd = &cobra.Command{
	Use:   "list-service-pipeline-outputs",
	Short: "Get a list of service pipeline Infrastructure as Code (IaC) outputs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listServicePipelineOutputsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listServicePipelineOutputsCmd).Standalone()

		proton_listServicePipelineOutputsCmd.Flags().String("deployment-id", "", "The ID of the deployment you want the outputs for.")
		proton_listServicePipelineOutputsCmd.Flags().String("next-token", "", "A token that indicates the location of the next output in the array of outputs, after the list of outputs that was previously requested.")
		proton_listServicePipelineOutputsCmd.Flags().String("service-name", "", "The name of the service whose pipeline's outputs you want.")
		proton_listServicePipelineOutputsCmd.MarkFlagRequired("service-name")
	})
	protonCmd.AddCommand(proton_listServicePipelineOutputsCmd)
}
