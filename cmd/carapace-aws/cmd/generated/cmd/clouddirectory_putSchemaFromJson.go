package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_putSchemaFromJsonCmd = &cobra.Command{
	Use:   "put-schema-from-json",
	Short: "Allows a schema to be updated using JSON upload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_putSchemaFromJsonCmd).Standalone()

	clouddirectory_putSchemaFromJsonCmd.Flags().String("document", "", "The replacement JSON schema.")
	clouddirectory_putSchemaFromJsonCmd.Flags().String("schema-arn", "", "The ARN of the schema to update.")
	clouddirectory_putSchemaFromJsonCmd.MarkFlagRequired("document")
	clouddirectory_putSchemaFromJsonCmd.MarkFlagRequired("schema-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_putSchemaFromJsonCmd)
}
