package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getServiceInstanceCmd = &cobra.Command{
	Use:   "get-service-instance",
	Short: "Get detailed data for a service instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getServiceInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_getServiceInstanceCmd).Standalone()

		proton_getServiceInstanceCmd.Flags().String("name", "", "The name of a service instance that you want to get the detailed data for.")
		proton_getServiceInstanceCmd.Flags().String("service-name", "", "The name of the service that you want the service instance input for.")
		proton_getServiceInstanceCmd.MarkFlagRequired("name")
		proton_getServiceInstanceCmd.MarkFlagRequired("service-name")
	})
	protonCmd.AddCommand(proton_getServiceInstanceCmd)
}
