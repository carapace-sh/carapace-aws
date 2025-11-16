package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_addFlowVpcInterfacesCmd = &cobra.Command{
	Use:   "add-flow-vpc-interfaces",
	Short: "Adds VPC interfaces to a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_addFlowVpcInterfacesCmd).Standalone()

	mediaconnect_addFlowVpcInterfacesCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to update.")
	mediaconnect_addFlowVpcInterfacesCmd.Flags().String("vpc-interfaces", "", "A list of VPC interfaces that you want to add to the flow.")
	mediaconnect_addFlowVpcInterfacesCmd.MarkFlagRequired("flow-arn")
	mediaconnect_addFlowVpcInterfacesCmd.MarkFlagRequired("vpc-interfaces")
	mediaconnectCmd.AddCommand(mediaconnect_addFlowVpcInterfacesCmd)
}
