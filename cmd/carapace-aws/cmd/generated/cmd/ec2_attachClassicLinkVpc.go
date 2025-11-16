package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_attachClassicLinkVpcCmd = &cobra.Command{
	Use:   "attach-classic-link-vpc",
	Short: "This action is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_attachClassicLinkVpcCmd).Standalone()

	ec2_attachClassicLinkVpcCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_attachClassicLinkVpcCmd.Flags().String("groups", "", "The IDs of the security groups.")
	ec2_attachClassicLinkVpcCmd.Flags().String("instance-id", "", "The ID of the EC2-Classic instance.")
	ec2_attachClassicLinkVpcCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_attachClassicLinkVpcCmd.Flags().String("vpc-id", "", "The ID of the ClassicLink-enabled VPC.")
	ec2_attachClassicLinkVpcCmd.MarkFlagRequired("groups")
	ec2_attachClassicLinkVpcCmd.MarkFlagRequired("instance-id")
	ec2_attachClassicLinkVpcCmd.Flag("no-dry-run").Hidden = true
	ec2_attachClassicLinkVpcCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_attachClassicLinkVpcCmd)
}
