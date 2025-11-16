package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_getClusterCredentialsWithIamCmd = &cobra.Command{
	Use:   "get-cluster-credentials-with-iam",
	Short: "Returns a database user name and temporary password with temporary authorization to log in to an Amazon Redshift database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_getClusterCredentialsWithIamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_getClusterCredentialsWithIamCmd).Standalone()

		redshift_getClusterCredentialsWithIamCmd.Flags().String("cluster-identifier", "", "The unique identifier of the cluster that contains the database for which you are requesting credentials.")
		redshift_getClusterCredentialsWithIamCmd.Flags().String("custom-domain-name", "", "The custom domain name for the IAM message cluster credentials.")
		redshift_getClusterCredentialsWithIamCmd.Flags().String("db-name", "", "The name of the database for which you are requesting credentials.")
		redshift_getClusterCredentialsWithIamCmd.Flags().String("duration-seconds", "", "The number of seconds until the returned temporary password expires.")
	})
	redshiftCmd.AddCommand(redshift_getClusterCredentialsWithIamCmd)
}
