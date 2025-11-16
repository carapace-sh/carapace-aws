package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listServiceInstanceOutputsCmd = &cobra.Command{
	Use:   "list-service-instance-outputs",
	Short: "Get a list service of instance Infrastructure as Code (IaC) outputs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listServiceInstanceOutputsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listServiceInstanceOutputsCmd).Standalone()

		proton_listServiceInstanceOutputsCmd.Flags().String("deployment-id", "", "The ID of the deployment whose outputs you want.")
		proton_listServiceInstanceOutputsCmd.Flags().String("next-token", "", "A token that indicates the location of the next output in the array of outputs, after the list of outputs that was previously requested.")
		proton_listServiceInstanceOutputsCmd.Flags().String("service-instance-name", "", "The name of the service instance whose outputs you want.")
		proton_listServiceInstanceOutputsCmd.Flags().String("service-name", "", "The name of the service that `serviceInstanceName` is associated to.")
		proton_listServiceInstanceOutputsCmd.MarkFlagRequired("service-instance-name")
		proton_listServiceInstanceOutputsCmd.MarkFlagRequired("service-name")
	})
	protonCmd.AddCommand(proton_listServiceInstanceOutputsCmd)
}
