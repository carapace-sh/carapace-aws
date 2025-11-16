package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeClusterParametersCmd = &cobra.Command{
	Use:   "describe-cluster-parameters",
	Short: "Returns a detailed list of parameters contained within the specified Amazon Redshift parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeClusterParametersCmd).Standalone()

	redshift_describeClusterParametersCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeClusterParametersCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_describeClusterParametersCmd.Flags().String("parameter-group-name", "", "The name of a cluster parameter group for which to return details.")
	redshift_describeClusterParametersCmd.Flags().String("source", "", "The parameter types to return.")
	redshift_describeClusterParametersCmd.MarkFlagRequired("parameter-group-name")
	redshiftCmd.AddCommand(redshift_describeClusterParametersCmd)
}
