package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_getClusterCredentialsCmd = &cobra.Command{
	Use:   "get-cluster-credentials",
	Short: "Returns a database user name and temporary password with temporary authorization to log on to an Amazon Redshift database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_getClusterCredentialsCmd).Standalone()

	redshift_getClusterCredentialsCmd.Flags().String("auto-create", "", "Create a database user with the name specified for the user named in `DbUser` if one does not exist.")
	redshift_getClusterCredentialsCmd.Flags().String("cluster-identifier", "", "The unique identifier of the cluster that contains the database for which you are requesting credentials.")
	redshift_getClusterCredentialsCmd.Flags().String("custom-domain-name", "", "The custom domain name for the cluster credentials.")
	redshift_getClusterCredentialsCmd.Flags().String("db-groups", "", "A list of the names of existing database groups that the user named in `DbUser` will join for the current session, in addition to any group memberships for an existing user.")
	redshift_getClusterCredentialsCmd.Flags().String("db-name", "", "The name of a database that `DbUser` is authorized to log on to.")
	redshift_getClusterCredentialsCmd.Flags().String("db-user", "", "The name of a database user.")
	redshift_getClusterCredentialsCmd.Flags().String("duration-seconds", "", "The number of seconds until the returned temporary password expires.")
	redshift_getClusterCredentialsCmd.MarkFlagRequired("db-user")
	redshiftCmd.AddCommand(redshift_getClusterCredentialsCmd)
}
