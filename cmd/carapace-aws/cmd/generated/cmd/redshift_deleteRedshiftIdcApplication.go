package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteRedshiftIdcApplicationCmd = &cobra.Command{
	Use:   "delete-redshift-idc-application",
	Short: "Deletes an Amazon Redshift IAM Identity Center application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteRedshiftIdcApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_deleteRedshiftIdcApplicationCmd).Standalone()

		redshift_deleteRedshiftIdcApplicationCmd.Flags().String("redshift-idc-application-arn", "", "The ARN for a deleted Amazon Redshift IAM Identity Center application.")
		redshift_deleteRedshiftIdcApplicationCmd.MarkFlagRequired("redshift-idc-application-arn")
	})
	redshiftCmd.AddCommand(redshift_deleteRedshiftIdcApplicationCmd)
}
