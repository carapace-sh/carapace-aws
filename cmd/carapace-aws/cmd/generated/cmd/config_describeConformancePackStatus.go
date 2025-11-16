package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeConformancePackStatusCmd = &cobra.Command{
	Use:   "describe-conformance-pack-status",
	Short: "Provides one or more conformance packs deployment status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeConformancePackStatusCmd).Standalone()

	config_describeConformancePackStatusCmd.Flags().String("conformance-pack-names", "", "Comma-separated list of conformance pack names.")
	config_describeConformancePackStatusCmd.Flags().String("limit", "", "The maximum number of conformance packs status returned on each page.")
	config_describeConformancePackStatusCmd.Flags().String("next-token", "", "The `nextToken` string returned in a previous request that you use to request the next page of results in a paginated response.")
	configCmd.AddCommand(config_describeConformancePackStatusCmd)
}
