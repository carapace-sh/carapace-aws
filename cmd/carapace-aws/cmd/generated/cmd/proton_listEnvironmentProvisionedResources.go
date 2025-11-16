package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listEnvironmentProvisionedResourcesCmd = &cobra.Command{
	Use:   "list-environment-provisioned-resources",
	Short: "List the provisioned resources for your environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listEnvironmentProvisionedResourcesCmd).Standalone()

	proton_listEnvironmentProvisionedResourcesCmd.Flags().String("environment-name", "", "The environment name.")
	proton_listEnvironmentProvisionedResourcesCmd.Flags().String("next-token", "", "A token that indicates the location of the next environment provisioned resource in the array of environment provisioned resources, after the list of environment provisioned resources that was previously requested.")
	proton_listEnvironmentProvisionedResourcesCmd.MarkFlagRequired("environment-name")
	protonCmd.AddCommand(proton_listEnvironmentProvisionedResourcesCmd)
}
