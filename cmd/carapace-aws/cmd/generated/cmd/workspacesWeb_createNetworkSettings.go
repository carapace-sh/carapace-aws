package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_createNetworkSettingsCmd = &cobra.Command{
	Use:   "create-network-settings",
	Short: "Creates a network settings resource that can be associated with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_createNetworkSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_createNetworkSettingsCmd).Standalone()

		workspacesWeb_createNetworkSettingsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		workspacesWeb_createNetworkSettingsCmd.Flags().String("security-group-ids", "", "One or more security groups used to control access from streaming instances to your VPC.")
		workspacesWeb_createNetworkSettingsCmd.Flags().String("subnet-ids", "", "The subnets in which network interfaces are created to connect streaming instances to your VPC.")
		workspacesWeb_createNetworkSettingsCmd.Flags().String("tags", "", "The tags to add to the network settings resource.")
		workspacesWeb_createNetworkSettingsCmd.Flags().String("vpc-id", "", "The VPC that streaming instances will connect to.")
		workspacesWeb_createNetworkSettingsCmd.MarkFlagRequired("security-group-ids")
		workspacesWeb_createNetworkSettingsCmd.MarkFlagRequired("subnet-ids")
		workspacesWeb_createNetworkSettingsCmd.MarkFlagRequired("vpc-id")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_createNetworkSettingsCmd)
}
