package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeClusterSubnetGroupsCmd = &cobra.Command{
	Use:   "describe-cluster-subnet-groups",
	Short: "Returns one or more cluster subnet group objects, which contain metadata about your cluster subnet groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeClusterSubnetGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeClusterSubnetGroupsCmd).Standalone()

		redshift_describeClusterSubnetGroupsCmd.Flags().String("cluster-subnet-group-name", "", "The name of the cluster subnet group for which information is requested.")
		redshift_describeClusterSubnetGroupsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
		redshift_describeClusterSubnetGroupsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
		redshift_describeClusterSubnetGroupsCmd.Flags().String("tag-keys", "", "A tag key or keys for which you want to return all matching cluster subnet groups that are associated with the specified key or keys.")
		redshift_describeClusterSubnetGroupsCmd.Flags().String("tag-values", "", "A tag value or values for which you want to return all matching cluster subnet groups that are associated with the specified tag value or values.")
	})
	redshiftCmd.AddCommand(redshift_describeClusterSubnetGroupsCmd)
}
