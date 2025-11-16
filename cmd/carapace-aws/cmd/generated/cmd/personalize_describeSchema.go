package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeSchemaCmd = &cobra.Command{
	Use:   "describe-schema",
	Short: "Describes a schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeSchemaCmd).Standalone()

	personalize_describeSchemaCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) of the schema to retrieve.")
	personalize_describeSchemaCmd.MarkFlagRequired("schema-arn")
	personalizeCmd.AddCommand(personalize_describeSchemaCmd)
}
