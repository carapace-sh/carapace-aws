package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRelationalDatabasesCmd = &cobra.Command{
	Use:   "get-relational-databases",
	Short: "Returns information about all of your databases in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRelationalDatabasesCmd).Standalone()

	lightsail_getRelationalDatabasesCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getRelationalDatabasesCmd)
}
