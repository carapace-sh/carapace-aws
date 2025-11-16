package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startTriggerCmd = &cobra.Command{
	Use:   "start-trigger",
	Short: "Starts an existing trigger.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startTriggerCmd).Standalone()

	glue_startTriggerCmd.Flags().String("name", "", "The name of the trigger to start.")
	glue_startTriggerCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_startTriggerCmd)
}
