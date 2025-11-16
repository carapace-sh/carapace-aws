package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpcBlockPublicAccessOptionsCmd = &cobra.Command{
	Use:   "modify-vpc-block-public-access-options",
	Short: "Modify VPC Block Public Access (BPA) options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpcBlockPublicAccessOptionsCmd).Standalone()

	ec2_modifyVpcBlockPublicAccessOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVpcBlockPublicAccessOptionsCmd.Flags().String("internet-gateway-block-mode", "", "The mode of VPC BPA.")
	ec2_modifyVpcBlockPublicAccessOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVpcBlockPublicAccessOptionsCmd.MarkFlagRequired("internet-gateway-block-mode")
	ec2_modifyVpcBlockPublicAccessOptionsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyVpcBlockPublicAccessOptionsCmd)
}
