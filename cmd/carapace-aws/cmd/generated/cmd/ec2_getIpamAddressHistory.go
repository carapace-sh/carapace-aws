package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getIpamAddressHistoryCmd = &cobra.Command{
	Use:   "get-ipam-address-history",
	Short: "Retrieve historical information about a CIDR within an IPAM scope.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getIpamAddressHistoryCmd).Standalone()

	ec2_getIpamAddressHistoryCmd.Flags().String("cidr", "", "The CIDR you want the history of.")
	ec2_getIpamAddressHistoryCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_getIpamAddressHistoryCmd.Flags().String("end-time", "", "The end of the time period for which you are looking for history.")
	ec2_getIpamAddressHistoryCmd.Flags().String("ipam-scope-id", "", "The ID of the IPAM scope that the CIDR is in.")
	ec2_getIpamAddressHistoryCmd.Flags().String("max-results", "", "The maximum number of historical results you would like returned per page.")
	ec2_getIpamAddressHistoryCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_getIpamAddressHistoryCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_getIpamAddressHistoryCmd.Flags().String("start-time", "", "The start of the time period for which you are looking for history.")
	ec2_getIpamAddressHistoryCmd.Flags().String("vpc-id", "", "The ID of the VPC you want your history records filtered by.")
	ec2_getIpamAddressHistoryCmd.MarkFlagRequired("cidr")
	ec2_getIpamAddressHistoryCmd.MarkFlagRequired("ipam-scope-id")
	ec2_getIpamAddressHistoryCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_getIpamAddressHistoryCmd)
}
