package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeUsageLimitsCmd = &cobra.Command{
	Use:   "describe-usage-limits",
	Short: "Shows usage limits on a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeUsageLimitsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeUsageLimitsCmd).Standalone()

		redshift_describeUsageLimitsCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster for which you want to describe usage limits.")
		redshift_describeUsageLimitsCmd.Flags().String("feature-type", "", "The feature type for which you want to describe usage limits.")
		redshift_describeUsageLimitsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
		redshift_describeUsageLimitsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
		redshift_describeUsageLimitsCmd.Flags().String("tag-keys", "", "A tag key or keys for which you want to return all matching usage limit objects that are associated with the specified key or keys.")
		redshift_describeUsageLimitsCmd.Flags().String("tag-values", "", "A tag value or values for which you want to return all matching usage limit objects that are associated with the specified tag value or values.")
		redshift_describeUsageLimitsCmd.Flags().String("usage-limit-id", "", "The identifier of the usage limit to describe.")
	})
	redshiftCmd.AddCommand(redshift_describeUsageLimitsCmd)
}
