package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getVpcAttachmentCmd = &cobra.Command{
	Use:   "get-vpc-attachment",
	Short: "Returns information about a VPC attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getVpcAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getVpcAttachmentCmd).Standalone()

		networkmanager_getVpcAttachmentCmd.Flags().String("attachment-id", "", "The ID of the attachment.")
		networkmanager_getVpcAttachmentCmd.MarkFlagRequired("attachment-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getVpcAttachmentCmd)
}
