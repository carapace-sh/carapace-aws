package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteListCmd = &cobra.Command{
	Use:   "delete-list",
	Short: "Deletes the list, provided it is not used in a rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_deleteListCmd).Standalone()

		frauddetector_deleteListCmd.Flags().String("name", "", "The name of the list to delete.")
		frauddetector_deleteListCmd.MarkFlagRequired("name")
	})
	frauddetectorCmd.AddCommand(frauddetector_deleteListCmd)
}
