package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_getPrivateGraphEndpointCmd = &cobra.Command{
	Use:   "get-private-graph-endpoint",
	Short: "Retrieves information about a specified private endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_getPrivateGraphEndpointCmd).Standalone()

	neptuneGraph_getPrivateGraphEndpointCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
	neptuneGraph_getPrivateGraphEndpointCmd.Flags().String("vpc-id", "", "The ID of the VPC where the private endpoint is located.")
	neptuneGraph_getPrivateGraphEndpointCmd.MarkFlagRequired("graph-identifier")
	neptuneGraph_getPrivateGraphEndpointCmd.MarkFlagRequired("vpc-id")
	neptuneGraphCmd.AddCommand(neptuneGraph_getPrivateGraphEndpointCmd)
}
