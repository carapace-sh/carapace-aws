package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_rebootRelationalDatabaseCmd = &cobra.Command{
	Use:   "reboot-relational-database",
	Short: "Restarts a specific database in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_rebootRelationalDatabaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_rebootRelationalDatabaseCmd).Standalone()

		lightsail_rebootRelationalDatabaseCmd.Flags().String("relational-database-name", "", "The name of your database to reboot.")
		lightsail_rebootRelationalDatabaseCmd.MarkFlagRequired("relational-database-name")
	})
	lightsailCmd.AddCommand(lightsail_rebootRelationalDatabaseCmd)
}
