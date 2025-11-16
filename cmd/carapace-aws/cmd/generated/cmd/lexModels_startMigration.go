package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_startMigrationCmd = &cobra.Command{
	Use:   "start-migration",
	Short: "Starts migrating a bot from Amazon Lex V1 to Amazon Lex V2.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_startMigrationCmd).Standalone()

	lexModels_startMigrationCmd.Flags().String("migration-strategy", "", "The strategy used to conduct the migration.")
	lexModels_startMigrationCmd.Flags().String("v1-bot-name", "", "The name of the Amazon Lex V1 bot that you are migrating to Amazon Lex V2.")
	lexModels_startMigrationCmd.Flags().String("v1-bot-version", "", "The version of the bot to migrate to Amazon Lex V2.")
	lexModels_startMigrationCmd.Flags().String("v2-bot-name", "", "The name of the Amazon Lex V2 bot that you are migrating the Amazon Lex V1 bot to.")
	lexModels_startMigrationCmd.Flags().String("v2-bot-role", "", "The IAM role that Amazon Lex uses to run the Amazon Lex V2 bot.")
	lexModels_startMigrationCmd.MarkFlagRequired("migration-strategy")
	lexModels_startMigrationCmd.MarkFlagRequired("v1-bot-name")
	lexModels_startMigrationCmd.MarkFlagRequired("v1-bot-version")
	lexModels_startMigrationCmd.MarkFlagRequired("v2-bot-name")
	lexModels_startMigrationCmd.MarkFlagRequired("v2-bot-role")
	lexModelsCmd.AddCommand(lexModels_startMigrationCmd)
}
