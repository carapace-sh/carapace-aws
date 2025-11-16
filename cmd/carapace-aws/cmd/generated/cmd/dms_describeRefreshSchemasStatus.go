package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeRefreshSchemasStatusCmd = &cobra.Command{
	Use:   "describe-refresh-schemas-status",
	Short: "Returns the status of the RefreshSchemas operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeRefreshSchemasStatusCmd).Standalone()

	dms_describeRefreshSchemasStatusCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.")
	dms_describeRefreshSchemasStatusCmd.MarkFlagRequired("endpoint-arn")
	dmsCmd.AddCommand(dms_describeRefreshSchemasStatusCmd)
}
