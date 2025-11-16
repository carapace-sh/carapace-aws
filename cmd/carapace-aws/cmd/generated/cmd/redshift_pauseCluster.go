package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_pauseClusterCmd = &cobra.Command{
	Use:   "pause-cluster",
	Short: "Pauses a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_pauseClusterCmd).Standalone()

	redshift_pauseClusterCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster to be paused.")
	redshift_pauseClusterCmd.MarkFlagRequired("cluster-identifier")
	redshiftCmd.AddCommand(redshift_pauseClusterCmd)
}
