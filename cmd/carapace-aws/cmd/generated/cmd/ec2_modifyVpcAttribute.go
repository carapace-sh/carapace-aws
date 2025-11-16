package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpcAttributeCmd = &cobra.Command{
	Use:   "modify-vpc-attribute",
	Short: "Modifies the specified attribute of the specified VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpcAttributeCmd).Standalone()

	ec2_modifyVpcAttributeCmd.Flags().String("enable-dns-hostnames", "", "Indicates whether the instances launched in the VPC get DNS hostnames.")
	ec2_modifyVpcAttributeCmd.Flags().String("enable-dns-support", "", "Indicates whether the DNS resolution is supported for the VPC.")
	ec2_modifyVpcAttributeCmd.Flags().String("enable-network-address-usage-metrics", "", "Indicates whether Network Address Usage metrics are enabled for your VPC.")
	ec2_modifyVpcAttributeCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
	ec2_modifyVpcAttributeCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_modifyVpcAttributeCmd)
}
