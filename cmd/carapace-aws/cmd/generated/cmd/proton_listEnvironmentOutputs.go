package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listEnvironmentOutputsCmd = &cobra.Command{
	Use:   "list-environment-outputs",
	Short: "List the infrastructure as code outputs for your environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listEnvironmentOutputsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listEnvironmentOutputsCmd).Standalone()

		proton_listEnvironmentOutputsCmd.Flags().String("deployment-id", "", "The ID of the deployment whose outputs you want.")
		proton_listEnvironmentOutputsCmd.Flags().String("environment-name", "", "The environment name.")
		proton_listEnvironmentOutputsCmd.Flags().String("next-token", "", "A token that indicates the location of the next environment output in the array of environment outputs, after the list of environment outputs that was previously requested.")
		proton_listEnvironmentOutputsCmd.MarkFlagRequired("environment-name")
	})
	protonCmd.AddCommand(proton_listEnvironmentOutputsCmd)
}
