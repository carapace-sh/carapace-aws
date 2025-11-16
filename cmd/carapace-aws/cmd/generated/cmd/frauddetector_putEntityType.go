package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_putEntityTypeCmd = &cobra.Command{
	Use:   "put-entity-type",
	Short: "Creates or updates an entity type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_putEntityTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_putEntityTypeCmd).Standalone()

		frauddetector_putEntityTypeCmd.Flags().String("description", "", "The description.")
		frauddetector_putEntityTypeCmd.Flags().String("name", "", "The name of the entity type.")
		frauddetector_putEntityTypeCmd.Flags().String("tags", "", "A collection of key and value pairs.")
		frauddetector_putEntityTypeCmd.MarkFlagRequired("name")
	})
	frauddetectorCmd.AddCommand(frauddetector_putEntityTypeCmd)
}
