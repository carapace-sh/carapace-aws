package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableVpcClassicLinkCmd = &cobra.Command{
	Use:   "enable-vpc-classic-link",
	Short: "This action is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableVpcClassicLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableVpcClassicLinkCmd).Standalone()

		ec2_enableVpcClassicLinkCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableVpcClassicLinkCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableVpcClassicLinkCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
		ec2_enableVpcClassicLinkCmd.Flag("no-dry-run").Hidden = true
		ec2_enableVpcClassicLinkCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_enableVpcClassicLinkCmd)
}
