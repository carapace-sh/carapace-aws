package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_refreshSchemasCmd = &cobra.Command{
	Use:   "refresh-schemas",
	Short: "Populates the schema for the specified endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_refreshSchemasCmd).Standalone()

	dms_refreshSchemasCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.")
	dms_refreshSchemasCmd.Flags().String("replication-instance-arn", "", "The Amazon Resource Name (ARN) of the replication instance.")
	dms_refreshSchemasCmd.MarkFlagRequired("endpoint-arn")
	dms_refreshSchemasCmd.MarkFlagRequired("replication-instance-arn")
	dmsCmd.AddCommand(dms_refreshSchemasCmd)
}
