package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getEntityTypesCmd = &cobra.Command{
	Use:   "get-entity-types",
	Short: "Gets all entity types or a specific entity type if a name is specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getEntityTypesCmd).Standalone()

	frauddetector_getEntityTypesCmd.Flags().String("max-results", "", "The maximum number of objects to return for the request.")
	frauddetector_getEntityTypesCmd.Flags().String("name", "", "The name.")
	frauddetector_getEntityTypesCmd.Flags().String("next-token", "", "The next token for the subsequent request.")
	frauddetectorCmd.AddCommand(frauddetector_getEntityTypesCmd)
}
