package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeSchemasCmd = &cobra.Command{
	Use:   "describe-schemas",
	Short: "Returns information about the schema for the specified endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeSchemasCmd).Standalone()

	dms_describeSchemasCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.")
	dms_describeSchemasCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeSchemasCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dms_describeSchemasCmd.MarkFlagRequired("endpoint-arn")
	dmsCmd.AddCommand(dms_describeSchemasCmd)
}
