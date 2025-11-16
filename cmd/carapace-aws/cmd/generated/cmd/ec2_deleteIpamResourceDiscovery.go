package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteIpamResourceDiscoveryCmd = &cobra.Command{
	Use:   "delete-ipam-resource-discovery",
	Short: "Deletes an IPAM resource discovery.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteIpamResourceDiscoveryCmd).Standalone()

	ec2_deleteIpamResourceDiscoveryCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deleteIpamResourceDiscoveryCmd.Flags().String("ipam-resource-discovery-id", "", "The IPAM resource discovery ID.")
	ec2_deleteIpamResourceDiscoveryCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deleteIpamResourceDiscoveryCmd.MarkFlagRequired("ipam-resource-discovery-id")
	ec2_deleteIpamResourceDiscoveryCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteIpamResourceDiscoveryCmd)
}
