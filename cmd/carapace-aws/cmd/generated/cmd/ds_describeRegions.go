package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeRegionsCmd = &cobra.Command{
	Use:   "describe-regions",
	Short: "Provides information about the Regions that are configured for multi-Region replication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeRegionsCmd).Standalone()

	ds_describeRegionsCmd.Flags().String("directory-id", "", "The identifier of the directory.")
	ds_describeRegionsCmd.Flags().String("next-token", "", "The `DescribeRegionsResult.NextToken` value from a previous call to [DescribeRegions]().")
	ds_describeRegionsCmd.Flags().String("region-name", "", "The name of the Region.")
	ds_describeRegionsCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_describeRegionsCmd)
}
