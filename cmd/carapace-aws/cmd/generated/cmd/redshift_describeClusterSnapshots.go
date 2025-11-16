package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeClusterSnapshotsCmd = &cobra.Command{
	Use:   "describe-cluster-snapshots",
	Short: "Returns one or more snapshot objects, which contain metadata about your cluster snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeClusterSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeClusterSnapshotsCmd).Standalone()

		redshift_describeClusterSnapshotsCmd.Flags().String("cluster-exists", "", "A value that indicates whether to return snapshots only for an existing cluster.")
		redshift_describeClusterSnapshotsCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster which generated the requested snapshots.")
		redshift_describeClusterSnapshotsCmd.Flags().String("end-time", "", "A time value that requests only snapshots created at or before the specified time.")
		redshift_describeClusterSnapshotsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
		redshift_describeClusterSnapshotsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
		redshift_describeClusterSnapshotsCmd.Flags().String("owner-account", "", "The Amazon Web Services account used to create or copy the snapshot.")
		redshift_describeClusterSnapshotsCmd.Flags().String("snapshot-arn", "", "The Amazon Resource Name (ARN) of the snapshot associated with the message to describe cluster snapshots.")
		redshift_describeClusterSnapshotsCmd.Flags().String("snapshot-identifier", "", "The snapshot identifier of the snapshot about which to return information.")
		redshift_describeClusterSnapshotsCmd.Flags().String("snapshot-type", "", "The type of snapshots for which you are requesting information.")
		redshift_describeClusterSnapshotsCmd.Flags().String("sorting-entities", "", "")
		redshift_describeClusterSnapshotsCmd.Flags().String("start-time", "", "A value that requests only snapshots created at or after the specified time.")
		redshift_describeClusterSnapshotsCmd.Flags().String("tag-keys", "", "A tag key or keys for which you want to return all matching cluster snapshots that are associated with the specified key or keys.")
		redshift_describeClusterSnapshotsCmd.Flags().String("tag-values", "", "A tag value or values for which you want to return all matching cluster snapshots that are associated with the specified tag value or values.")
	})
	redshiftCmd.AddCommand(redshift_describeClusterSnapshotsCmd)
}
