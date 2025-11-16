package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_deleteSourceNetworkCmd = &cobra.Command{
	Use:   "delete-source-network",
	Short: "Delete Source Network resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_deleteSourceNetworkCmd).Standalone()

	drs_deleteSourceNetworkCmd.Flags().String("source-network-id", "", "ID of the Source Network to delete.")
	drs_deleteSourceNetworkCmd.MarkFlagRequired("source-network-id")
	drsCmd.AddCommand(drs_deleteSourceNetworkCmd)
}
