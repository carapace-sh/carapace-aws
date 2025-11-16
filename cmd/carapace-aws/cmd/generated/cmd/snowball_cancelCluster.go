package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_cancelClusterCmd = &cobra.Command{
	Use:   "cancel-cluster",
	Short: "Cancels a cluster job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_cancelClusterCmd).Standalone()

	snowball_cancelClusterCmd.Flags().String("cluster-id", "", "The 39-character ID for the cluster that you want to cancel, for example `CID123e4567-e89b-12d3-a456-426655440000`.")
	snowball_cancelClusterCmd.MarkFlagRequired("cluster-id")
	snowballCmd.AddCommand(snowball_cancelClusterCmd)
}
