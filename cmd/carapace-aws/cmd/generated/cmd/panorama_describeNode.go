package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_describeNodeCmd = &cobra.Command{
	Use:   "describe-node",
	Short: "Returns information about a node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_describeNodeCmd).Standalone()

	panorama_describeNodeCmd.Flags().String("node-id", "", "The node's ID.")
	panorama_describeNodeCmd.Flags().String("owner-account", "", "The account ID of the node's owner.")
	panorama_describeNodeCmd.MarkFlagRequired("node-id")
	panoramaCmd.AddCommand(panorama_describeNodeCmd)
}
