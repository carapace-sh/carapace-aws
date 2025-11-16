package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_applySchemaCmd = &cobra.Command{
	Use:   "apply-schema",
	Short: "Copies the input published schema, at the specified version, into the [Directory]() with the same name and version as that of the published schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_applySchemaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_applySchemaCmd).Standalone()

		clouddirectory_applySchemaCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() into which the schema is copied.")
		clouddirectory_applySchemaCmd.Flags().String("published-schema-arn", "", "Published schema Amazon Resource Name (ARN) that needs to be copied.")
		clouddirectory_applySchemaCmd.MarkFlagRequired("directory-arn")
		clouddirectory_applySchemaCmd.MarkFlagRequired("published-schema-arn")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_applySchemaCmd)
}
