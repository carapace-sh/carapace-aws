package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_createSchemaCmd = &cobra.Command{
	Use:   "create-schema",
	Short: "Creates a new schema in a development state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_createSchemaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_createSchemaCmd).Standalone()

		clouddirectory_createSchemaCmd.Flags().String("name", "", "The name that is associated with the schema.")
		clouddirectory_createSchemaCmd.MarkFlagRequired("name")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_createSchemaCmd)
}
