package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_startDbclusterCmd = &cobra.Command{
	Use:   "start-dbcluster",
	Short: "Restarts the stopped cluster that is specified by `DBClusterIdentifier`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_startDbclusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_startDbclusterCmd).Standalone()

		docdb_startDbclusterCmd.Flags().String("dbcluster-identifier", "", "The identifier of the cluster to restart.")
		docdb_startDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	})
	docdbCmd.AddCommand(docdb_startDbclusterCmd)
}
