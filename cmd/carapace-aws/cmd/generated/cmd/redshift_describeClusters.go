package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeClustersCmd = &cobra.Command{
	Use:   "describe-clusters",
	Short: "Returns properties of provisioned clusters including general cluster properties, cluster database properties, maintenance and backup properties, and security and access properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeClustersCmd).Standalone()

		redshift_describeClustersCmd.Flags().String("cluster-identifier", "", "The unique identifier of a cluster whose properties you are requesting.")
		redshift_describeClustersCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
		redshift_describeClustersCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
		redshift_describeClustersCmd.Flags().String("tag-keys", "", "A tag key or keys for which you want to return all matching clusters that are associated with the specified key or keys.")
		redshift_describeClustersCmd.Flags().String("tag-values", "", "A tag value or values for which you want to return all matching clusters that are associated with the specified tag value or values.")
	})
	redshiftCmd.AddCommand(redshift_describeClustersCmd)
}
