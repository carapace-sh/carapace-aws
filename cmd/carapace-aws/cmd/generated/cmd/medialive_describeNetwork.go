package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeNetworkCmd = &cobra.Command{
	Use:   "describe-network",
	Short: "Get details about a Network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeNetworkCmd).Standalone()

	medialive_describeNetworkCmd.Flags().String("network-id", "", "The ID of the network.")
	medialive_describeNetworkCmd.MarkFlagRequired("network-id")
	medialiveCmd.AddCommand(medialive_describeNetworkCmd)
}
