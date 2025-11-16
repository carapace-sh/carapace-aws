package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listServicePipelineProvisionedResourcesCmd = &cobra.Command{
	Use:   "list-service-pipeline-provisioned-resources",
	Short: "List provisioned resources for a service and pipeline with details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listServicePipelineProvisionedResourcesCmd).Standalone()

	proton_listServicePipelineProvisionedResourcesCmd.Flags().String("next-token", "", "A token that indicates the location of the next provisioned resource in the array of provisioned resources, after the list of provisioned resources that was previously requested.")
	proton_listServicePipelineProvisionedResourcesCmd.Flags().String("service-name", "", "The name of the service whose pipeline's provisioned resources you want.")
	proton_listServicePipelineProvisionedResourcesCmd.MarkFlagRequired("service-name")
	protonCmd.AddCommand(proton_listServicePipelineProvisionedResourcesCmd)
}
