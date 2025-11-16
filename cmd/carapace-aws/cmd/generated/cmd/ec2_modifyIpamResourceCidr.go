package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyIpamResourceCidrCmd = &cobra.Command{
	Use:   "modify-ipam-resource-cidr",
	Short: "Modify a resource CIDR.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyIpamResourceCidrCmd).Standalone()

	ec2_modifyIpamResourceCidrCmd.Flags().String("current-ipam-scope-id", "", "The ID of the current scope that the resource CIDR is in.")
	ec2_modifyIpamResourceCidrCmd.Flags().String("destination-ipam-scope-id", "", "The ID of the scope you want to transfer the resource CIDR to.")
	ec2_modifyIpamResourceCidrCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_modifyIpamResourceCidrCmd.Flags().Bool("monitored", false, "Determines if the resource is monitored by IPAM.")
	ec2_modifyIpamResourceCidrCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_modifyIpamResourceCidrCmd.Flags().Bool("no-monitored", false, "Determines if the resource is monitored by IPAM.")
	ec2_modifyIpamResourceCidrCmd.Flags().String("resource-cidr", "", "The CIDR of the resource you want to modify.")
	ec2_modifyIpamResourceCidrCmd.Flags().String("resource-id", "", "The ID of the resource you want to modify.")
	ec2_modifyIpamResourceCidrCmd.Flags().String("resource-region", "", "The Amazon Web Services Region of the resource you want to modify.")
	ec2_modifyIpamResourceCidrCmd.MarkFlagRequired("current-ipam-scope-id")
	ec2_modifyIpamResourceCidrCmd.MarkFlagRequired("monitored")
	ec2_modifyIpamResourceCidrCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyIpamResourceCidrCmd.Flag("no-monitored").Hidden = true
	ec2_modifyIpamResourceCidrCmd.MarkFlagRequired("no-monitored")
	ec2_modifyIpamResourceCidrCmd.MarkFlagRequired("resource-cidr")
	ec2_modifyIpamResourceCidrCmd.MarkFlagRequired("resource-id")
	ec2_modifyIpamResourceCidrCmd.MarkFlagRequired("resource-region")
	ec2Cmd.AddCommand(ec2_modifyIpamResourceCidrCmd)
}
