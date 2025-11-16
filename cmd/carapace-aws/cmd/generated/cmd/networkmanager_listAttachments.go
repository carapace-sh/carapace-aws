package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_listAttachmentsCmd = &cobra.Command{
	Use:   "list-attachments",
	Short: "Returns a list of core network attachments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_listAttachmentsCmd).Standalone()

	networkmanager_listAttachmentsCmd.Flags().String("attachment-type", "", "The type of attachment.")
	networkmanager_listAttachmentsCmd.Flags().String("core-network-id", "", "The ID of a core network.")
	networkmanager_listAttachmentsCmd.Flags().String("edge-location", "", "The Region where the edge is located.")
	networkmanager_listAttachmentsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	networkmanager_listAttachmentsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	networkmanager_listAttachmentsCmd.Flags().String("state", "", "The state of the attachment.")
	networkmanagerCmd.AddCommand(networkmanager_listAttachmentsCmd)
}
