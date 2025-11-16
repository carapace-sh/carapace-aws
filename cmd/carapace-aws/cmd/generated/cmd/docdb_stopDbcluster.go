package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_stopDbclusterCmd = &cobra.Command{
	Use:   "stop-dbcluster",
	Short: "Stops the running cluster that is specified by `DBClusterIdentifier`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_stopDbclusterCmd).Standalone()

	docdb_stopDbclusterCmd.Flags().String("dbcluster-identifier", "", "The identifier of the cluster to stop.")
	docdb_stopDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	docdbCmd.AddCommand(docdb_stopDbclusterCmd)
}
