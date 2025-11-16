package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_updateSchemaCmd = &cobra.Command{
	Use:   "update-schema",
	Short: "Updates the schema name with a new name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_updateSchemaCmd).Standalone()

	clouddirectory_updateSchemaCmd.Flags().String("name", "", "The name of the schema.")
	clouddirectory_updateSchemaCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) of the development schema.")
	clouddirectory_updateSchemaCmd.MarkFlagRequired("name")
	clouddirectory_updateSchemaCmd.MarkFlagRequired("schema-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_updateSchemaCmd)
}
