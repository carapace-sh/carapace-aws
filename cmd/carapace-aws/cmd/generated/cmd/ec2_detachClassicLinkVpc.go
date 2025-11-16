package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_detachClassicLinkVpcCmd = &cobra.Command{
	Use:   "detach-classic-link-vpc",
	Short: "This action is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_detachClassicLinkVpcCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_detachClassicLinkVpcCmd).Standalone()

		ec2_detachClassicLinkVpcCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_detachClassicLinkVpcCmd.Flags().String("instance-id", "", "The ID of the instance to unlink from the VPC.")
		ec2_detachClassicLinkVpcCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_detachClassicLinkVpcCmd.Flags().String("vpc-id", "", "The ID of the VPC to which the instance is linked.")
		ec2_detachClassicLinkVpcCmd.MarkFlagRequired("instance-id")
		ec2_detachClassicLinkVpcCmd.Flag("no-dry-run").Hidden = true
		ec2_detachClassicLinkVpcCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_detachClassicLinkVpcCmd)
}
