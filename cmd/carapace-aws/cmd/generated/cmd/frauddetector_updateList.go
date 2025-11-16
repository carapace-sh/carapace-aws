package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_updateListCmd = &cobra.Command{
	Use:   "update-list",
	Short: "Updates a list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_updateListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_updateListCmd).Standalone()

		frauddetector_updateListCmd.Flags().String("description", "", "The new description.")
		frauddetector_updateListCmd.Flags().String("elements", "", "One or more list elements to add or replace.")
		frauddetector_updateListCmd.Flags().String("name", "", "The name of the list to update.")
		frauddetector_updateListCmd.Flags().String("update-mode", "", "The update mode (type).")
		frauddetector_updateListCmd.Flags().String("variable-type", "", "The variable type you want to assign to the list.")
		frauddetector_updateListCmd.MarkFlagRequired("name")
	})
	frauddetectorCmd.AddCommand(frauddetector_updateListCmd)
}
