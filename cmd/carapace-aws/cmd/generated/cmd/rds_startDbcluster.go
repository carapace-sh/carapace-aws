package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_startDbclusterCmd = &cobra.Command{
	Use:   "start-dbcluster",
	Short: "Starts an Amazon Aurora DB cluster that was stopped using the Amazon Web Services console, the stop-db-cluster CLI command, or the `StopDBCluster` operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_startDbclusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_startDbclusterCmd).Standalone()

		rds_startDbclusterCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier of the Amazon Aurora DB cluster to be started.")
		rds_startDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	})
	rdsCmd.AddCommand(rds_startDbclusterCmd)
}
