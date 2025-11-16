package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_updateVpcAttachmentCmd = &cobra.Command{
	Use:   "update-vpc-attachment",
	Short: "Updates a VPC attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_updateVpcAttachmentCmd).Standalone()

	networkmanager_updateVpcAttachmentCmd.Flags().String("add-subnet-arns", "", "Adds a subnet ARN to the VPC attachment.")
	networkmanager_updateVpcAttachmentCmd.Flags().String("attachment-id", "", "The ID of the attachment.")
	networkmanager_updateVpcAttachmentCmd.Flags().String("options", "", "Additional options for updating the VPC attachment.")
	networkmanager_updateVpcAttachmentCmd.Flags().String("remove-subnet-arns", "", "Removes a subnet ARN from the attachment.")
	networkmanager_updateVpcAttachmentCmd.MarkFlagRequired("attachment-id")
	networkmanagerCmd.AddCommand(networkmanager_updateVpcAttachmentCmd)
}
