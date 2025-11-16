package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getMigrationCmd = &cobra.Command{
	Use:   "get-migration",
	Short: "Provides details about an ongoing or complete migration from an Amazon Lex V1 bot to an Amazon Lex V2 bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getMigrationCmd).Standalone()

	lexModels_getMigrationCmd.Flags().String("migration-id", "", "The unique identifier of the migration to view.")
	lexModels_getMigrationCmd.MarkFlagRequired("migration-id")
	lexModelsCmd.AddCommand(lexModels_getMigrationCmd)
}
