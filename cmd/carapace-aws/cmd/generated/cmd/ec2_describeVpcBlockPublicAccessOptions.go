package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcBlockPublicAccessOptionsCmd = &cobra.Command{
	Use:   "describe-vpc-block-public-access-options",
	Short: "Describe VPC Block Public Access (BPA) options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcBlockPublicAccessOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVpcBlockPublicAccessOptionsCmd).Standalone()

		ec2_describeVpcBlockPublicAccessOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcBlockPublicAccessOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcBlockPublicAccessOptionsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVpcBlockPublicAccessOptionsCmd)
}
