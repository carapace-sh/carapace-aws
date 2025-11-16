package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_startDbclusterCmd = &cobra.Command{
	Use:   "start-dbcluster",
	Short: "Starts an Amazon Neptune DB cluster that was stopped using the Amazon console, the Amazon CLI stop-db-cluster command, or the StopDBCluster API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_startDbclusterCmd).Standalone()

	neptune_startDbclusterCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier of the Neptune DB cluster to be started.")
	neptune_startDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	neptuneCmd.AddCommand(neptune_startDbclusterCmd)
}
