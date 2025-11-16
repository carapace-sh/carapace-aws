package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRelationalDatabaseBundlesCmd = &cobra.Command{
	Use:   "get-relational-database-bundles",
	Short: "Returns the list of bundles that are available in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRelationalDatabaseBundlesCmd).Standalone()

	lightsail_getRelationalDatabaseBundlesCmd.Flags().String("include-inactive", "", "A Boolean value that indicates whether to include inactive (unavailable) bundles in the response of your request.")
	lightsail_getRelationalDatabaseBundlesCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getRelationalDatabaseBundlesCmd)
}
