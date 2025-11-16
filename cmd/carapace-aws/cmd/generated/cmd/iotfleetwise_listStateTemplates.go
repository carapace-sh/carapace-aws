package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listStateTemplatesCmd = &cobra.Command{
	Use:   "list-state-templates",
	Short: "Lists information about created state templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listStateTemplatesCmd).Standalone()

	iotfleetwise_listStateTemplatesCmd.Flags().String("list-response-scope", "", "When you set the `listResponseScope` parameter to `METADATA_ONLY`, the list response includes: state template ID, Amazon Resource Name (ARN), creation time, and last modification time.")
	iotfleetwise_listStateTemplatesCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
	iotfleetwise_listStateTemplatesCmd.Flags().String("next-token", "", "The token to retrieve the next set of results, or `null` if there are no more results.")
	iotfleetwiseCmd.AddCommand(iotfleetwise_listStateTemplatesCmd)
}
