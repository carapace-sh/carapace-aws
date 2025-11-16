package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listServiceInstanceProvisionedResourcesCmd = &cobra.Command{
	Use:   "list-service-instance-provisioned-resources",
	Short: "List provisioned resources for a service instance with details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listServiceInstanceProvisionedResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listServiceInstanceProvisionedResourcesCmd).Standalone()

		proton_listServiceInstanceProvisionedResourcesCmd.Flags().String("next-token", "", "A token that indicates the location of the next provisioned resource in the array of provisioned resources, after the list of provisioned resources that was previously requested.")
		proton_listServiceInstanceProvisionedResourcesCmd.Flags().String("service-instance-name", "", "The name of the service instance whose provisioned resources you want.")
		proton_listServiceInstanceProvisionedResourcesCmd.Flags().String("service-name", "", "The name of the service that `serviceInstanceName` is associated to.")
		proton_listServiceInstanceProvisionedResourcesCmd.MarkFlagRequired("service-instance-name")
		proton_listServiceInstanceProvisionedResourcesCmd.MarkFlagRequired("service-name")
	})
	protonCmd.AddCommand(proton_listServiceInstanceProvisionedResourcesCmd)
}
