package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createVpcAttachmentCmd = &cobra.Command{
	Use:   "create-vpc-attachment",
	Short: "Creates a VPC attachment on an edge location of a core network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createVpcAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_createVpcAttachmentCmd).Standalone()

		networkmanager_createVpcAttachmentCmd.Flags().String("client-token", "", "The client token associated with the request.")
		networkmanager_createVpcAttachmentCmd.Flags().String("core-network-id", "", "The ID of a core network for the VPC attachment.")
		networkmanager_createVpcAttachmentCmd.Flags().String("options", "", "Options for the VPC attachment.")
		networkmanager_createVpcAttachmentCmd.Flags().String("subnet-arns", "", "The subnet ARN of the VPC attachment.")
		networkmanager_createVpcAttachmentCmd.Flags().String("tags", "", "The key-value tags associated with the request.")
		networkmanager_createVpcAttachmentCmd.Flags().String("vpc-arn", "", "The ARN of the VPC.")
		networkmanager_createVpcAttachmentCmd.MarkFlagRequired("core-network-id")
		networkmanager_createVpcAttachmentCmd.MarkFlagRequired("subnet-arns")
		networkmanager_createVpcAttachmentCmd.MarkFlagRequired("vpc-arn")
	})
	networkmanagerCmd.AddCommand(networkmanager_createVpcAttachmentCmd)
}
