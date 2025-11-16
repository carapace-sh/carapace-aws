package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listResourceGatewaysCmd = &cobra.Command{
	Use:   "list-resource-gateways",
	Short: "Lists the resource gateways that you own or that were shared with you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listResourceGatewaysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_listResourceGatewaysCmd).Standalone()

		vpcLattice_listResourceGatewaysCmd.Flags().String("max-results", "", "The maximum page size.")
		vpcLattice_listResourceGatewaysCmd.Flags().String("next-token", "", "If there are additional results, a pagination token for the next page of results.")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_listResourceGatewaysCmd)
}
