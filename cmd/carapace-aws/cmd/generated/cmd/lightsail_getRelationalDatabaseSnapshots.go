package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRelationalDatabaseSnapshotsCmd = &cobra.Command{
	Use:   "get-relational-database-snapshots",
	Short: "Returns information about all of your database snapshots in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRelationalDatabaseSnapshotsCmd).Standalone()

	lightsail_getRelationalDatabaseSnapshotsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getRelationalDatabaseSnapshotsCmd)
}
