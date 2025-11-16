package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deletePublicIpv4PoolCmd = &cobra.Command{
	Use:   "delete-public-ipv4-pool",
	Short: "Delete a public IPv4 pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deletePublicIpv4PoolCmd).Standalone()

	ec2_deletePublicIpv4PoolCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deletePublicIpv4PoolCmd.Flags().String("network-border-group", "", "The Availability Zone (AZ) or Local Zone (LZ) network border group that the resource that the IP address is assigned to is in.")
	ec2_deletePublicIpv4PoolCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deletePublicIpv4PoolCmd.Flags().String("pool-id", "", "The ID of the public IPv4 pool you want to delete.")
	ec2_deletePublicIpv4PoolCmd.Flag("no-dry-run").Hidden = true
	ec2_deletePublicIpv4PoolCmd.MarkFlagRequired("pool-id")
	ec2Cmd.AddCommand(ec2_deletePublicIpv4PoolCmd)
}
