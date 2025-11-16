package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeClusterParameterGroupsCmd = &cobra.Command{
	Use:   "describe-cluster-parameter-groups",
	Short: "Returns a list of Amazon Redshift parameter groups, including parameter groups you created and the default parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeClusterParameterGroupsCmd).Standalone()

	redshift_describeClusterParameterGroupsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeClusterParameterGroupsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_describeClusterParameterGroupsCmd.Flags().String("parameter-group-name", "", "The name of a specific parameter group for which to return details.")
	redshift_describeClusterParameterGroupsCmd.Flags().String("tag-keys", "", "A tag key or keys for which you want to return all matching cluster parameter groups that are associated with the specified key or keys.")
	redshift_describeClusterParameterGroupsCmd.Flags().String("tag-values", "", "A tag value or values for which you want to return all matching cluster parameter groups that are associated with the specified tag value or values.")
	redshiftCmd.AddCommand(redshift_describeClusterParameterGroupsCmd)
}
