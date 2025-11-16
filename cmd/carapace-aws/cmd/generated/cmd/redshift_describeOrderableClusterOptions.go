package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeOrderableClusterOptionsCmd = &cobra.Command{
	Use:   "describe-orderable-cluster-options",
	Short: "Returns a list of orderable cluster options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeOrderableClusterOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeOrderableClusterOptionsCmd).Standalone()

		redshift_describeOrderableClusterOptionsCmd.Flags().String("cluster-version", "", "The version filter value.")
		redshift_describeOrderableClusterOptionsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
		redshift_describeOrderableClusterOptionsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
		redshift_describeOrderableClusterOptionsCmd.Flags().String("node-type", "", "The node type filter value.")
	})
	redshiftCmd.AddCommand(redshift_describeOrderableClusterOptionsCmd)
}
