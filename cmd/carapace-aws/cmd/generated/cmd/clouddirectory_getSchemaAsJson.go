package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_getSchemaAsJsonCmd = &cobra.Command{
	Use:   "get-schema-as-json",
	Short: "Retrieves a JSON representation of the schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_getSchemaAsJsonCmd).Standalone()

	clouddirectory_getSchemaAsJsonCmd.Flags().String("schema-arn", "", "The ARN of the schema to retrieve.")
	clouddirectory_getSchemaAsJsonCmd.MarkFlagRequired("schema-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_getSchemaAsJsonCmd)
}
