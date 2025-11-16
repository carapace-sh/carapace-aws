package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_stopTriggerCmd = &cobra.Command{
	Use:   "stop-trigger",
	Short: "Stops a specified trigger.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_stopTriggerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_stopTriggerCmd).Standalone()

		glue_stopTriggerCmd.Flags().String("name", "", "The name of the trigger to stop.")
		glue_stopTriggerCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_stopTriggerCmd)
}
