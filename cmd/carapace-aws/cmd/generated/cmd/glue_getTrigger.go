package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getTriggerCmd = &cobra.Command{
	Use:   "get-trigger",
	Short: "Retrieves the definition of a trigger.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getTriggerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getTriggerCmd).Standalone()

		glue_getTriggerCmd.Flags().String("name", "", "The name of the trigger to retrieve.")
		glue_getTriggerCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_getTriggerCmd)
}
