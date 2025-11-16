package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeClusterVersionsCmd = &cobra.Command{
	Use:   "describe-cluster-versions",
	Short: "Returns descriptions of the available Amazon Redshift cluster versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeClusterVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeClusterVersionsCmd).Standalone()

		redshift_describeClusterVersionsCmd.Flags().String("cluster-parameter-group-family", "", "The name of a specific cluster parameter group family to return details for.")
		redshift_describeClusterVersionsCmd.Flags().String("cluster-version", "", "The specific cluster version to return.")
		redshift_describeClusterVersionsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
		redshift_describeClusterVersionsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	})
	redshiftCmd.AddCommand(redshift_describeClusterVersionsCmd)
}
