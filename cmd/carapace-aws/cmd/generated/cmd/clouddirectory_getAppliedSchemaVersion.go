package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_getAppliedSchemaVersionCmd = &cobra.Command{
	Use:   "get-applied-schema-version",
	Short: "Returns current applied schema version ARN, including the minor version in use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_getAppliedSchemaVersionCmd).Standalone()

	clouddirectory_getAppliedSchemaVersionCmd.Flags().String("schema-arn", "", "The ARN of the applied schema.")
	clouddirectory_getAppliedSchemaVersionCmd.MarkFlagRequired("schema-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_getAppliedSchemaVersionCmd)
}
