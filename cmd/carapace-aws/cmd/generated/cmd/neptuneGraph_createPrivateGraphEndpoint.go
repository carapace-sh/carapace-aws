package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_createPrivateGraphEndpointCmd = &cobra.Command{
	Use:   "create-private-graph-endpoint",
	Short: "Create a private graph endpoint to allow private access from to the graph from within a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_createPrivateGraphEndpointCmd).Standalone()

	neptuneGraph_createPrivateGraphEndpointCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
	neptuneGraph_createPrivateGraphEndpointCmd.Flags().String("subnet-ids", "", "Subnets in which private graph endpoint ENIs are created.")
	neptuneGraph_createPrivateGraphEndpointCmd.Flags().String("vpc-id", "", "The VPC in which the private graph endpoint needs to be created.")
	neptuneGraph_createPrivateGraphEndpointCmd.Flags().String("vpc-security-group-ids", "", "Security groups to be attached to the private graph endpoint..")
	neptuneGraph_createPrivateGraphEndpointCmd.MarkFlagRequired("graph-identifier")
	neptuneGraphCmd.AddCommand(neptuneGraph_createPrivateGraphEndpointCmd)
}
