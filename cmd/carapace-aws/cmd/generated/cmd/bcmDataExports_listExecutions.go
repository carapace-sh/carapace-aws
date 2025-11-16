package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExports_listExecutionsCmd = &cobra.Command{
	Use:   "list-executions",
	Short: "Lists the historical executions for the export.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExports_listExecutionsCmd).Standalone()

	bcmDataExports_listExecutionsCmd.Flags().String("export-arn", "", "The Amazon Resource Name (ARN) for this export.")
	bcmDataExports_listExecutionsCmd.Flags().String("max-results", "", "The maximum number of objects that are returned for the request.")
	bcmDataExports_listExecutionsCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	bcmDataExports_listExecutionsCmd.MarkFlagRequired("export-arn")
	bcmDataExportsCmd.AddCommand(bcmDataExports_listExecutionsCmd)
}
