package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateTriggerCmd = &cobra.Command{
	Use:   "update-trigger",
	Short: "Updates a trigger definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateTriggerCmd).Standalone()

	glue_updateTriggerCmd.Flags().String("name", "", "The name of the trigger to update.")
	glue_updateTriggerCmd.Flags().String("trigger-update", "", "The new values with which to update the trigger.")
	glue_updateTriggerCmd.MarkFlagRequired("name")
	glue_updateTriggerCmd.MarkFlagRequired("trigger-update")
	glueCmd.AddCommand(glue_updateTriggerCmd)
}
