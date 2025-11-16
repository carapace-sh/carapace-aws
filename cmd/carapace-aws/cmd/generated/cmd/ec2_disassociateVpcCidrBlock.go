package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateVpcCidrBlockCmd = &cobra.Command{
	Use:   "disassociate-vpc-cidr-block",
	Short: "Disassociates a CIDR block from a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateVpcCidrBlockCmd).Standalone()

	ec2_disassociateVpcCidrBlockCmd.Flags().String("association-id", "", "The association ID for the CIDR block.")
	ec2_disassociateVpcCidrBlockCmd.MarkFlagRequired("association-id")
	ec2Cmd.AddCommand(ec2_disassociateVpcCidrBlockCmd)
}
