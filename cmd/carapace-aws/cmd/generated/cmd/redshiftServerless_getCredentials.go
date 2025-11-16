package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getCredentialsCmd = &cobra.Command{
	Use:   "get-credentials",
	Short: "Returns a database user name and temporary password with temporary authorization to log in to Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getCredentialsCmd).Standalone()

	redshiftServerless_getCredentialsCmd.Flags().String("custom-domain-name", "", "The custom domain name associated with the workgroup.")
	redshiftServerless_getCredentialsCmd.Flags().String("db-name", "", "The name of the database to get temporary authorization to log on to.")
	redshiftServerless_getCredentialsCmd.Flags().String("duration-seconds", "", "The number of seconds until the returned temporary password expires.")
	redshiftServerless_getCredentialsCmd.Flags().String("workgroup-name", "", "The name of the workgroup associated with the database.")
	redshiftServerlessCmd.AddCommand(redshiftServerless_getCredentialsCmd)
}
