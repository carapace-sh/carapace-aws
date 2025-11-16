package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_deleteDbclusterCmd = &cobra.Command{
	Use:   "delete-dbcluster",
	Short: "Deletes a previously provisioned cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_deleteDbclusterCmd).Standalone()

	docdb_deleteDbclusterCmd.Flags().String("dbcluster-identifier", "", "The cluster identifier for the cluster to be deleted.")
	docdb_deleteDbclusterCmd.Flags().String("final-dbsnapshot-identifier", "", "The cluster snapshot identifier of the new cluster snapshot created when `SkipFinalSnapshot` is set to `false`.")
	docdb_deleteDbclusterCmd.Flags().Bool("no-skip-final-snapshot", false, "Determines whether a final cluster snapshot is created before the cluster is deleted.")
	docdb_deleteDbclusterCmd.Flags().Bool("skip-final-snapshot", false, "Determines whether a final cluster snapshot is created before the cluster is deleted.")
	docdb_deleteDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	docdb_deleteDbclusterCmd.Flag("no-skip-final-snapshot").Hidden = true
	docdbCmd.AddCommand(docdb_deleteDbclusterCmd)
}
