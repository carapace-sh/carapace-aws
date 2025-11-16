package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableVpcClassicLinkDnsSupportCmd = &cobra.Command{
	Use:   "enable-vpc-classic-link-dns-support",
	Short: "This action is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableVpcClassicLinkDnsSupportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableVpcClassicLinkDnsSupportCmd).Standalone()

		ec2_enableVpcClassicLinkDnsSupportCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
	})
	ec2Cmd.AddCommand(ec2_enableVpcClassicLinkDnsSupportCmd)
}
