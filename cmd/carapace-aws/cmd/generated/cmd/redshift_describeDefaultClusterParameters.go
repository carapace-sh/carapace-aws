package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeDefaultClusterParametersCmd = &cobra.Command{
	Use:   "describe-default-cluster-parameters",
	Short: "Returns a list of parameter settings for the specified parameter group family.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeDefaultClusterParametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeDefaultClusterParametersCmd).Standalone()

		redshift_describeDefaultClusterParametersCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
		redshift_describeDefaultClusterParametersCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
		redshift_describeDefaultClusterParametersCmd.Flags().String("parameter-group-family", "", "The name of the cluster parameter group family.")
		redshift_describeDefaultClusterParametersCmd.MarkFlagRequired("parameter-group-family")
	})
	redshiftCmd.AddCommand(redshift_describeDefaultClusterParametersCmd)
}
