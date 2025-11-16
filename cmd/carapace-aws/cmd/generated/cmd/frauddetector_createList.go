package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_createListCmd = &cobra.Command{
	Use:   "create-list",
	Short: "Creates a list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_createListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_createListCmd).Standalone()

		frauddetector_createListCmd.Flags().String("description", "", "The description of the list.")
		frauddetector_createListCmd.Flags().String("elements", "", "The names of the elements, if providing.")
		frauddetector_createListCmd.Flags().String("name", "", "The name of the list.")
		frauddetector_createListCmd.Flags().String("tags", "", "A collection of the key and value pairs.")
		frauddetector_createListCmd.Flags().String("variable-type", "", "The variable type of the list.")
		frauddetector_createListCmd.MarkFlagRequired("name")
	})
	frauddetectorCmd.AddCommand(frauddetector_createListCmd)
}
