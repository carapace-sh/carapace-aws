package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Deletes a previously provisioned cluster without its final snapshot being created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_deleteClusterCmd).Standalone()

		redshift_deleteClusterCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster to be deleted.")
		redshift_deleteClusterCmd.Flags().String("final-cluster-snapshot-identifier", "", "The identifier of the final snapshot that is to be created immediately before deleting the cluster.")
		redshift_deleteClusterCmd.Flags().String("final-cluster-snapshot-retention-period", "", "The number of days that a manual snapshot is retained.")
		redshift_deleteClusterCmd.Flags().Bool("no-skip-final-cluster-snapshot", false, "Determines whether a final snapshot of the cluster is created before Amazon Redshift deletes the cluster.")
		redshift_deleteClusterCmd.Flags().Bool("skip-final-cluster-snapshot", false, "Determines whether a final snapshot of the cluster is created before Amazon Redshift deletes the cluster.")
		redshift_deleteClusterCmd.MarkFlagRequired("cluster-identifier")
		redshift_deleteClusterCmd.Flag("no-skip-final-cluster-snapshot").Hidden = true
	})
	redshiftCmd.AddCommand(redshift_deleteClusterCmd)
}
