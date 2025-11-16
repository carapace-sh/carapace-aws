package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteNetworkCmd = &cobra.Command{
	Use:   "delete-network",
	Short: "Delete a Network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteNetworkCmd).Standalone()

	medialive_deleteNetworkCmd.Flags().String("network-id", "", "The ID of the network.")
	medialive_deleteNetworkCmd.MarkFlagRequired("network-id")
	medialiveCmd.AddCommand(medialive_deleteNetworkCmd)
}
