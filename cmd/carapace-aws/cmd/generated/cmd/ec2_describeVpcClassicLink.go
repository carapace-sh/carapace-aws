package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcClassicLinkCmd = &cobra.Command{
	Use:   "describe-vpc-classic-link",
	Short: "This action is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcClassicLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVpcClassicLinkCmd).Standalone()

		ec2_describeVpcClassicLinkCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcClassicLinkCmd.Flags().String("filters", "", "The filters.")
		ec2_describeVpcClassicLinkCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcClassicLinkCmd.Flags().String("vpc-ids", "", "The VPCs for which you want to describe the ClassicLink status.")
		ec2_describeVpcClassicLinkCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVpcClassicLinkCmd)
}
