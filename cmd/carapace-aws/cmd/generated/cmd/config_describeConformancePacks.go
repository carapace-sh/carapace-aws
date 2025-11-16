package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeConformancePacksCmd = &cobra.Command{
	Use:   "describe-conformance-packs",
	Short: "Returns a list of one or more conformance packs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeConformancePacksCmd).Standalone()

	config_describeConformancePacksCmd.Flags().String("conformance-pack-names", "", "Comma-separated list of conformance pack names for which you want details.")
	config_describeConformancePacksCmd.Flags().String("limit", "", "The maximum number of conformance packs returned on each page.")
	config_describeConformancePacksCmd.Flags().String("next-token", "", "The `nextToken` string returned in a previous request that you use to request the next page of results in a paginated response.")
	configCmd.AddCommand(config_describeConformancePacksCmd)
}
