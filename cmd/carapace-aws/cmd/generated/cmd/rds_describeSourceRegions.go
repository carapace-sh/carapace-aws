package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeSourceRegionsCmd = &cobra.Command{
	Use:   "describe-source-regions",
	Short: "Returns a list of the source Amazon Web Services Regions where the current Amazon Web Services Region can create a read replica, copy a DB snapshot from, or replicate automated backups from.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeSourceRegionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeSourceRegionsCmd).Standalone()

		rds_describeSourceRegionsCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
		rds_describeSourceRegionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeSourceRegions` request.")
		rds_describeSourceRegionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeSourceRegionsCmd.Flags().String("region-name", "", "The source Amazon Web Services Region name.")
	})
	rdsCmd.AddCommand(rds_describeSourceRegionsCmd)
}
