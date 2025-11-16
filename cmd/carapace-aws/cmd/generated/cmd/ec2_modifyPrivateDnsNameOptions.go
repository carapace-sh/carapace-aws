package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyPrivateDnsNameOptionsCmd = &cobra.Command{
	Use:   "modify-private-dns-name-options",
	Short: "Modifies the options for instance hostnames for the specified instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyPrivateDnsNameOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyPrivateDnsNameOptionsCmd).Standalone()

		ec2_modifyPrivateDnsNameOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyPrivateDnsNameOptionsCmd.Flags().Bool("enable-resource-name-dns-aaaarecord", false, "Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.")
		ec2_modifyPrivateDnsNameOptionsCmd.Flags().Bool("enable-resource-name-dns-arecord", false, "Indicates whether to respond to DNS queries for instance hostnames with DNS A records.")
		ec2_modifyPrivateDnsNameOptionsCmd.Flags().String("instance-id", "", "The ID of the instance.")
		ec2_modifyPrivateDnsNameOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyPrivateDnsNameOptionsCmd.Flags().Bool("no-enable-resource-name-dns-aaaarecord", false, "Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.")
		ec2_modifyPrivateDnsNameOptionsCmd.Flags().Bool("no-enable-resource-name-dns-arecord", false, "Indicates whether to respond to DNS queries for instance hostnames with DNS A records.")
		ec2_modifyPrivateDnsNameOptionsCmd.Flags().String("private-dns-hostname-type", "", "The type of hostname for EC2 instances.")
		ec2_modifyPrivateDnsNameOptionsCmd.MarkFlagRequired("instance-id")
		ec2_modifyPrivateDnsNameOptionsCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyPrivateDnsNameOptionsCmd.Flag("no-enable-resource-name-dns-aaaarecord").Hidden = true
		ec2_modifyPrivateDnsNameOptionsCmd.Flag("no-enable-resource-name-dns-arecord").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyPrivateDnsNameOptionsCmd)
}
