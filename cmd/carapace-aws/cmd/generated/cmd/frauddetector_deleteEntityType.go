package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteEntityTypeCmd = &cobra.Command{
	Use:   "delete-entity-type",
	Short: "Deletes an entity type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteEntityTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_deleteEntityTypeCmd).Standalone()

		frauddetector_deleteEntityTypeCmd.Flags().String("name", "", "The name of the entity type to delete.")
		frauddetector_deleteEntityTypeCmd.MarkFlagRequired("name")
	})
	frauddetectorCmd.AddCommand(frauddetector_deleteEntityTypeCmd)
}
