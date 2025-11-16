package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_updateMemberCmd = &cobra.Command{
	Use:   "update-member",
	Short: "Updates a member configuration with new parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_updateMemberCmd).Standalone()

	managedblockchain_updateMemberCmd.Flags().String("log-publishing-configuration", "", "Configuration properties for publishing to Amazon CloudWatch Logs.")
	managedblockchain_updateMemberCmd.Flags().String("member-id", "", "The unique identifier of the member.")
	managedblockchain_updateMemberCmd.Flags().String("network-id", "", "The unique identifier of the Managed Blockchain network to which the member belongs.")
	managedblockchain_updateMemberCmd.MarkFlagRequired("member-id")
	managedblockchain_updateMemberCmd.MarkFlagRequired("network-id")
	managedblockchainCmd.AddCommand(managedblockchain_updateMemberCmd)
}
