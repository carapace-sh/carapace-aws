package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listComponentProvisionedResourcesCmd = &cobra.Command{
	Use:   "list-component-provisioned-resources",
	Short: "List provisioned resources for a component with details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listComponentProvisionedResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listComponentProvisionedResourcesCmd).Standalone()

		proton_listComponentProvisionedResourcesCmd.Flags().String("component-name", "", "The name of the component whose provisioned resources you want.")
		proton_listComponentProvisionedResourcesCmd.Flags().String("next-token", "", "A token that indicates the location of the next provisioned resource in the array of provisioned resources, after the list of provisioned resources that was previously requested.")
		proton_listComponentProvisionedResourcesCmd.MarkFlagRequired("component-name")
	})
	protonCmd.AddCommand(proton_listComponentProvisionedResourcesCmd)
}
