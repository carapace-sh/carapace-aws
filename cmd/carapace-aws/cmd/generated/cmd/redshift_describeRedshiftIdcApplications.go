package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeRedshiftIdcApplicationsCmd = &cobra.Command{
	Use:   "describe-redshift-idc-applications",
	Short: "Lists the Amazon Redshift IAM Identity Center applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeRedshiftIdcApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeRedshiftIdcApplicationsCmd).Standalone()

		redshift_describeRedshiftIdcApplicationsCmd.Flags().String("marker", "", "A value that indicates the starting point for the next set of response records in a subsequent request.")
		redshift_describeRedshiftIdcApplicationsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
		redshift_describeRedshiftIdcApplicationsCmd.Flags().String("redshift-idc-application-arn", "", "The ARN for the Redshift application that integrates with IAM Identity Center.")
	})
	redshiftCmd.AddCommand(redshift_describeRedshiftIdcApplicationsCmd)
}
