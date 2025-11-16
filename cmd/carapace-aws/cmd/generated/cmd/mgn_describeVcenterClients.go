package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_describeVcenterClientsCmd = &cobra.Command{
	Use:   "describe-vcenter-clients",
	Short: "Returns a list of the installed vCenter clients.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_describeVcenterClientsCmd).Standalone()

	mgn_describeVcenterClientsCmd.Flags().String("max-results", "", "Maximum results to be returned in DescribeVcenterClients.")
	mgn_describeVcenterClientsCmd.Flags().String("next-token", "", "Next pagination token to be provided for DescribeVcenterClients.")
	mgnCmd.AddCommand(mgn_describeVcenterClientsCmd)
}
