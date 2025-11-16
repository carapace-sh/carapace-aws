package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getSchemaCmd = &cobra.Command{
	Use:   "get-schema",
	Short: "Describes the specified schema in detail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getSchemaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getSchemaCmd).Standalone()

		glue_getSchemaCmd.Flags().String("schema-id", "", "This is a wrapper structure to contain schema identity fields.")
		glue_getSchemaCmd.MarkFlagRequired("schema-id")
	})
	glueCmd.AddCommand(glue_getSchemaCmd)
}
