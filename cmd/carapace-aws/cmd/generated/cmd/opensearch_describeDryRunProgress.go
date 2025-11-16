package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeDryRunProgressCmd = &cobra.Command{
	Use:   "describe-dry-run-progress",
	Short: "Describes the progress of a pre-update dry run analysis on an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeDryRunProgressCmd).Standalone()

	opensearch_describeDryRunProgressCmd.Flags().String("domain-name", "", "The name of the domain.")
	opensearch_describeDryRunProgressCmd.Flags().String("dry-run-id", "", "The unique identifier of the dry run.")
	opensearch_describeDryRunProgressCmd.Flags().Bool("load-dry-run-config", false, "Whether to include the configuration of the dry run in the response.")
	opensearch_describeDryRunProgressCmd.Flags().Bool("no-load-dry-run-config", false, "Whether to include the configuration of the dry run in the response.")
	opensearch_describeDryRunProgressCmd.MarkFlagRequired("domain-name")
	opensearch_describeDryRunProgressCmd.Flag("no-load-dry-run-config").Hidden = true
	opensearchCmd.AddCommand(opensearch_describeDryRunProgressCmd)
}
