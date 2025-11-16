package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getTemporaryGlueTableCredentialsCmd = &cobra.Command{
	Use:   "get-temporary-glue-table-credentials",
	Short: "Allows a caller in a secure environment to assume a role with permission to access Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getTemporaryGlueTableCredentialsCmd).Standalone()

	lakeformation_getTemporaryGlueTableCredentialsCmd.Flags().String("audit-context", "", "A structure representing context to access a resource (column names, query ID, etc).")
	lakeformation_getTemporaryGlueTableCredentialsCmd.Flags().String("duration-seconds", "", "The time period, between 900 and 21,600 seconds, for the timeout of the temporary credentials.")
	lakeformation_getTemporaryGlueTableCredentialsCmd.Flags().String("permissions", "", "Filters the request based on the user having been granted a list of specified permissions on the requested resource(s).")
	lakeformation_getTemporaryGlueTableCredentialsCmd.Flags().String("query-session-context", "", "A structure used as a protocol between query engines and Lake Formation or Glue.")
	lakeformation_getTemporaryGlueTableCredentialsCmd.Flags().String("s3-path", "", "The Amazon S3 path for the table.")
	lakeformation_getTemporaryGlueTableCredentialsCmd.Flags().String("supported-permission-types", "", "A list of supported permission types for the table.")
	lakeformation_getTemporaryGlueTableCredentialsCmd.Flags().String("table-arn", "", "The ARN identifying a table in the Data Catalog for the temporary credentials request.")
	lakeformation_getTemporaryGlueTableCredentialsCmd.MarkFlagRequired("table-arn")
	lakeformationCmd.AddCommand(lakeformation_getTemporaryGlueTableCredentialsCmd)
}
