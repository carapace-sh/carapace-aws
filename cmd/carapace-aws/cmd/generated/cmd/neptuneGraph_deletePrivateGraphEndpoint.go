package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_deletePrivateGraphEndpointCmd = &cobra.Command{
	Use:   "delete-private-graph-endpoint",
	Short: "Deletes a private graph endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_deletePrivateGraphEndpointCmd).Standalone()

	neptuneGraph_deletePrivateGraphEndpointCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
	neptuneGraph_deletePrivateGraphEndpointCmd.Flags().String("vpc-id", "", "The ID of the VPC where the private endpoint is located.")
	neptuneGraph_deletePrivateGraphEndpointCmd.MarkFlagRequired("graph-identifier")
	neptuneGraph_deletePrivateGraphEndpointCmd.MarkFlagRequired("vpc-id")
	neptuneGraphCmd.AddCommand(neptuneGraph_deletePrivateGraphEndpointCmd)
}
