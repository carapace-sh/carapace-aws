package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_removeFlowVpcInterfaceCmd = &cobra.Command{
	Use:   "remove-flow-vpc-interface",
	Short: "Removes a VPC Interface from an existing flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_removeFlowVpcInterfaceCmd).Standalone()

	mediaconnect_removeFlowVpcInterfaceCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to remove a VPC interface from.")
	mediaconnect_removeFlowVpcInterfaceCmd.Flags().String("vpc-interface-name", "", "The name of the VPC interface that you want to remove.")
	mediaconnect_removeFlowVpcInterfaceCmd.MarkFlagRequired("flow-arn")
	mediaconnect_removeFlowVpcInterfaceCmd.MarkFlagRequired("vpc-interface-name")
	mediaconnectCmd.AddCommand(mediaconnect_removeFlowVpcInterfaceCmd)
}
