package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createIpamResourceDiscoveryCmd = &cobra.Command{
	Use:   "create-ipam-resource-discovery",
	Short: "Creates an IPAM resource discovery.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createIpamResourceDiscoveryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createIpamResourceDiscoveryCmd).Standalone()

		ec2_createIpamResourceDiscoveryCmd.Flags().String("client-token", "", "A client token for the IPAM resource discovery.")
		ec2_createIpamResourceDiscoveryCmd.Flags().String("description", "", "A description for the IPAM resource discovery.")
		ec2_createIpamResourceDiscoveryCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_createIpamResourceDiscoveryCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_createIpamResourceDiscoveryCmd.Flags().String("operating-regions", "", "Operating Regions for the IPAM resource discovery.")
		ec2_createIpamResourceDiscoveryCmd.Flags().String("tag-specifications", "", "Tag specifications for the IPAM resource discovery.")
		ec2_createIpamResourceDiscoveryCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createIpamResourceDiscoveryCmd)
}
