package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableVpcClassicLinkDnsSupportCmd = &cobra.Command{
	Use:   "disable-vpc-classic-link-dns-support",
	Short: "This action is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableVpcClassicLinkDnsSupportCmd).Standalone()

	ec2_disableVpcClassicLinkDnsSupportCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
	ec2Cmd.AddCommand(ec2_disableVpcClassicLinkDnsSupportCmd)
}
