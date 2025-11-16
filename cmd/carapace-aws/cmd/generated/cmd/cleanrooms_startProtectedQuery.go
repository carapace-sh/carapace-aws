package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_startProtectedQueryCmd = &cobra.Command{
	Use:   "start-protected-query",
	Short: "Creates a protected query that is started by Clean Rooms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_startProtectedQueryCmd).Standalone()

	cleanrooms_startProtectedQueryCmd.Flags().String("compute-configuration", "", "The compute configuration for the protected query.")
	cleanrooms_startProtectedQueryCmd.Flags().String("membership-identifier", "", "A unique identifier for the membership to run this query against.")
	cleanrooms_startProtectedQueryCmd.Flags().String("result-configuration", "", "The details needed to write the query results.")
	cleanrooms_startProtectedQueryCmd.Flags().String("sql-parameters", "", "The protected SQL query parameters.")
	cleanrooms_startProtectedQueryCmd.Flags().String("type", "", "The type of the protected query to be started.")
	cleanrooms_startProtectedQueryCmd.MarkFlagRequired("membership-identifier")
	cleanrooms_startProtectedQueryCmd.MarkFlagRequired("sql-parameters")
	cleanrooms_startProtectedQueryCmd.MarkFlagRequired("type")
	cleanroomsCmd.AddCommand(cleanrooms_startProtectedQueryCmd)
}
