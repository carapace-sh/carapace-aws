package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeClusterDbRevisionsCmd = &cobra.Command{
	Use:   "describe-cluster-db-revisions",
	Short: "Returns an array of `ClusterDbRevision` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeClusterDbRevisionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeClusterDbRevisionsCmd).Standalone()

		redshift_describeClusterDbRevisionsCmd.Flags().String("cluster-identifier", "", "A unique identifier for a cluster whose `ClusterDbRevisions` you are requesting.")
		redshift_describeClusterDbRevisionsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point for returning a set of response records.")
		redshift_describeClusterDbRevisionsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	})
	redshiftCmd.AddCommand(redshift_describeClusterDbRevisionsCmd)
}
