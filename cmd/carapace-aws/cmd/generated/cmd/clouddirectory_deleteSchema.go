package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_deleteSchemaCmd = &cobra.Command{
	Use:   "delete-schema",
	Short: "Deletes a given schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_deleteSchemaCmd).Standalone()

	clouddirectory_deleteSchemaCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) of the development schema.")
	clouddirectory_deleteSchemaCmd.MarkFlagRequired("schema-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_deleteSchemaCmd)
}
