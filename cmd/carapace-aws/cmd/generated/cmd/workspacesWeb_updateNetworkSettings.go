package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_updateNetworkSettingsCmd = &cobra.Command{
	Use:   "update-network-settings",
	Short: "Updates network settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_updateNetworkSettingsCmd).Standalone()

	workspacesWeb_updateNetworkSettingsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesWeb_updateNetworkSettingsCmd.Flags().String("network-settings-arn", "", "The ARN of the network settings.")
	workspacesWeb_updateNetworkSettingsCmd.Flags().String("security-group-ids", "", "One or more security groups used to control access from streaming instances to your VPC.")
	workspacesWeb_updateNetworkSettingsCmd.Flags().String("subnet-ids", "", "The subnets in which network interfaces are created to connect streaming instances to your VPC.")
	workspacesWeb_updateNetworkSettingsCmd.Flags().String("vpc-id", "", "The VPC that streaming instances will connect to.")
	workspacesWeb_updateNetworkSettingsCmd.MarkFlagRequired("network-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_updateNetworkSettingsCmd)
}
