package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_createListenerCmd = &cobra.Command{
	Use:   "create-listener",
	Short: "Creates a listener for a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_createListenerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_createListenerCmd).Standalone()

		vpcLattice_createListenerCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		vpcLattice_createListenerCmd.Flags().String("default-action", "", "The action for the default rule.")
		vpcLattice_createListenerCmd.Flags().String("name", "", "The name of the listener.")
		vpcLattice_createListenerCmd.Flags().String("port", "", "The listener port.")
		vpcLattice_createListenerCmd.Flags().String("protocol", "", "The listener protocol.")
		vpcLattice_createListenerCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
		vpcLattice_createListenerCmd.Flags().String("tags", "", "The tags for the listener.")
		vpcLattice_createListenerCmd.MarkFlagRequired("default-action")
		vpcLattice_createListenerCmd.MarkFlagRequired("name")
		vpcLattice_createListenerCmd.MarkFlagRequired("protocol")
		vpcLattice_createListenerCmd.MarkFlagRequired("service-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_createListenerCmd)
}
