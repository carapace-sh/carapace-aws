package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_publishSchemaCmd = &cobra.Command{
	Use:   "publish-schema",
	Short: "Publishes a development schema with a major version and a recommended minor version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_publishSchemaCmd).Standalone()

	clouddirectory_publishSchemaCmd.Flags().String("development-schema-arn", "", "The Amazon Resource Name (ARN) that is associated with the development schema.")
	clouddirectory_publishSchemaCmd.Flags().String("minor-version", "", "The minor version under which the schema will be published.")
	clouddirectory_publishSchemaCmd.Flags().String("name", "", "The new name under which the schema will be published.")
	clouddirectory_publishSchemaCmd.Flags().String("version", "", "The major version under which the schema will be published.")
	clouddirectory_publishSchemaCmd.MarkFlagRequired("development-schema-arn")
	clouddirectory_publishSchemaCmd.MarkFlagRequired("version")
	clouddirectoryCmd.AddCommand(clouddirectory_publishSchemaCmd)
}
