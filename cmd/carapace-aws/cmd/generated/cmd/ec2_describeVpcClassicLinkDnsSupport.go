package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcClassicLinkDnsSupportCmd = &cobra.Command{
	Use:   "describe-vpc-classic-link-dns-support",
	Short: "This action is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcClassicLinkDnsSupportCmd).Standalone()

	ec2_describeVpcClassicLinkDnsSupportCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeVpcClassicLinkDnsSupportCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeVpcClassicLinkDnsSupportCmd.Flags().String("vpc-ids", "", "The IDs of the VPCs.")
	ec2Cmd.AddCommand(ec2_describeVpcClassicLinkDnsSupportCmd)
}
