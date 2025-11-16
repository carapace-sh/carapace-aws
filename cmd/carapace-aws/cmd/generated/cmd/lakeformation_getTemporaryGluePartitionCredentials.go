package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getTemporaryGluePartitionCredentialsCmd = &cobra.Command{
	Use:   "get-temporary-glue-partition-credentials",
	Short: "This API is identical to `GetTemporaryTableCredentials` except that this is used when the target Data Catalog resource is of type Partition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getTemporaryGluePartitionCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_getTemporaryGluePartitionCredentialsCmd).Standalone()

		lakeformation_getTemporaryGluePartitionCredentialsCmd.Flags().String("audit-context", "", "A structure representing context to access a resource (column names, query ID, etc).")
		lakeformation_getTemporaryGluePartitionCredentialsCmd.Flags().String("duration-seconds", "", "The time period, between 900 and 21,600 seconds, for the timeout of the temporary credentials.")
		lakeformation_getTemporaryGluePartitionCredentialsCmd.Flags().String("partition", "", "A list of partition values identifying a single partition.")
		lakeformation_getTemporaryGluePartitionCredentialsCmd.Flags().String("permissions", "", "Filters the request based on the user having been granted a list of specified permissions on the requested resource(s).")
		lakeformation_getTemporaryGluePartitionCredentialsCmd.Flags().String("supported-permission-types", "", "A list of supported permission types for the partition.")
		lakeformation_getTemporaryGluePartitionCredentialsCmd.Flags().String("table-arn", "", "The ARN of the partitions' table.")
		lakeformation_getTemporaryGluePartitionCredentialsCmd.MarkFlagRequired("partition")
		lakeformation_getTemporaryGluePartitionCredentialsCmd.MarkFlagRequired("table-arn")
	})
	lakeformationCmd.AddCommand(lakeformation_getTemporaryGluePartitionCredentialsCmd)
}
