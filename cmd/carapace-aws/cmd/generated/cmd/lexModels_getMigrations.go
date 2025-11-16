package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getMigrationsCmd = &cobra.Command{
	Use:   "get-migrations",
	Short: "Gets a list of migrations between Amazon Lex V1 and Amazon Lex V2.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getMigrationsCmd).Standalone()

	lexModels_getMigrationsCmd.Flags().String("max-results", "", "The maximum number of migrations to return in the response.")
	lexModels_getMigrationsCmd.Flags().String("migration-status-equals", "", "Filters the list to contain only migrations in the specified state.")
	lexModels_getMigrationsCmd.Flags().String("next-token", "", "A pagination token that fetches the next page of migrations.")
	lexModels_getMigrationsCmd.Flags().String("sort-by-attribute", "", "The field to sort the list of migrations by.")
	lexModels_getMigrationsCmd.Flags().String("sort-by-order", "", "The order so sort the list.")
	lexModels_getMigrationsCmd.Flags().String("v1-bot-name-contains", "", "Filters the list to contain only bots whose name contains the specified string.")
	lexModelsCmd.AddCommand(lexModels_getMigrationsCmd)
}
