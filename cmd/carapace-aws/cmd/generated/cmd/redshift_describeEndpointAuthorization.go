package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeEndpointAuthorizationCmd = &cobra.Command{
	Use:   "describe-endpoint-authorization",
	Short: "Describes an endpoint authorization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeEndpointAuthorizationCmd).Standalone()

	redshift_describeEndpointAuthorizationCmd.Flags().String("account", "", "The Amazon Web Services account ID of either the cluster owner (grantor) or grantee.")
	redshift_describeEndpointAuthorizationCmd.Flags().String("cluster-identifier", "", "The cluster identifier of the cluster to access.")
	redshift_describeEndpointAuthorizationCmd.Flags().String("grantee", "", "Indicates whether to check authorization from a grantor or grantee point of view.")
	redshift_describeEndpointAuthorizationCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeEndpointAuthorization` request.")
	redshift_describeEndpointAuthorizationCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	redshiftCmd.AddCommand(redshift_describeEndpointAuthorizationCmd)
}
