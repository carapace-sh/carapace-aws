package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbclustersCmd = &cobra.Command{
	Use:   "describe-dbclusters",
	Short: "Describes existing Amazon Aurora DB clusters and Multi-AZ DB clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbclustersCmd).Standalone()

	rds_describeDbclustersCmd.Flags().String("dbcluster-identifier", "", "The user-supplied DB cluster identifier or the Amazon Resource Name (ARN) of the DB cluster.")
	rds_describeDbclustersCmd.Flags().String("filters", "", "A filter that specifies one or more DB clusters to describe.")
	rds_describeDbclustersCmd.Flags().Bool("include-shared", false, "Specifies whether the output includes information about clusters shared from other Amazon Web Services accounts.")
	rds_describeDbclustersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBClusters` request.")
	rds_describeDbclustersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rds_describeDbclustersCmd.Flags().Bool("no-include-shared", false, "Specifies whether the output includes information about clusters shared from other Amazon Web Services accounts.")
	rds_describeDbclustersCmd.Flag("no-include-shared").Hidden = true
	rdsCmd.AddCommand(rds_describeDbclustersCmd)
}
