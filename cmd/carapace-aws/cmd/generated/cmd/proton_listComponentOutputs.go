package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listComponentOutputsCmd = &cobra.Command{
	Use:   "list-component-outputs",
	Short: "Get a list of component Infrastructure as Code (IaC) outputs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listComponentOutputsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listComponentOutputsCmd).Standalone()

		proton_listComponentOutputsCmd.Flags().String("component-name", "", "The name of the component whose outputs you want.")
		proton_listComponentOutputsCmd.Flags().String("deployment-id", "", "The ID of the deployment whose outputs you want.")
		proton_listComponentOutputsCmd.Flags().String("next-token", "", "A token that indicates the location of the next output in the array of outputs, after the list of outputs that was previously requested.")
		proton_listComponentOutputsCmd.MarkFlagRequired("component-name")
	})
	protonCmd.AddCommand(proton_listComponentOutputsCmd)
}
