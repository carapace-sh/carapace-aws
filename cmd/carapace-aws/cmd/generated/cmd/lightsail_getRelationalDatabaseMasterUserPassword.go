package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRelationalDatabaseMasterUserPasswordCmd = &cobra.Command{
	Use:   "get-relational-database-master-user-password",
	Short: "Returns the current, previous, or pending versions of the master user password for a Lightsail database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRelationalDatabaseMasterUserPasswordCmd).Standalone()

	lightsail_getRelationalDatabaseMasterUserPasswordCmd.Flags().String("password-version", "", "The password version to return.")
	lightsail_getRelationalDatabaseMasterUserPasswordCmd.Flags().String("relational-database-name", "", "The name of your database for which to get the master user password.")
	lightsail_getRelationalDatabaseMasterUserPasswordCmd.MarkFlagRequired("relational-database-name")
	lightsailCmd.AddCommand(lightsail_getRelationalDatabaseMasterUserPasswordCmd)
}
