package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateSubnetCidrBlockCmd = &cobra.Command{
	Use:   "disassociate-subnet-cidr-block",
	Short: "Disassociates a CIDR block from a subnet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateSubnetCidrBlockCmd).Standalone()

	ec2_disassociateSubnetCidrBlockCmd.Flags().String("association-id", "", "The association ID for the CIDR block.")
	ec2_disassociateSubnetCidrBlockCmd.MarkFlagRequired("association-id")
	ec2Cmd.AddCommand(ec2_disassociateSubnetCidrBlockCmd)
}
