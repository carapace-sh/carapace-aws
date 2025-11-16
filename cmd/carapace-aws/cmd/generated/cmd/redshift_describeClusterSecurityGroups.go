package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeClusterSecurityGroupsCmd = &cobra.Command{
	Use:   "describe-cluster-security-groups",
	Short: "Returns information about Amazon Redshift security groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeClusterSecurityGroupsCmd).Standalone()

	redshift_describeClusterSecurityGroupsCmd.Flags().String("cluster-security-group-name", "", "The name of a cluster security group for which you are requesting details.")
	redshift_describeClusterSecurityGroupsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeClusterSecurityGroupsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_describeClusterSecurityGroupsCmd.Flags().String("tag-keys", "", "A tag key or keys for which you want to return all matching cluster security groups that are associated with the specified key or keys.")
	redshift_describeClusterSecurityGroupsCmd.Flags().String("tag-values", "", "A tag value or values for which you want to return all matching cluster security groups that are associated with the specified tag value or values.")
	redshiftCmd.AddCommand(redshift_describeClusterSecurityGroupsCmd)
}
