package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteTriggerCmd = &cobra.Command{
	Use:   "delete-trigger",
	Short: "Deletes a specified trigger.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteTriggerCmd).Standalone()

	glue_deleteTriggerCmd.Flags().String("name", "", "The name of the trigger to delete.")
	glue_deleteTriggerCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_deleteTriggerCmd)
}
