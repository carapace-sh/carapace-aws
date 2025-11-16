package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyRedshiftIdcApplicationCmd = &cobra.Command{
	Use:   "modify-redshift-idc-application",
	Short: "Changes an existing Amazon Redshift IAM Identity Center application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyRedshiftIdcApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_modifyRedshiftIdcApplicationCmd).Standalone()

		redshift_modifyRedshiftIdcApplicationCmd.Flags().String("authorized-token-issuer-list", "", "The authorized token issuer list for the Amazon Redshift IAM Identity Center application to change.")
		redshift_modifyRedshiftIdcApplicationCmd.Flags().String("iam-role-arn", "", "The IAM role ARN associated with the Amazon Redshift IAM Identity Center application to change.")
		redshift_modifyRedshiftIdcApplicationCmd.Flags().String("idc-display-name", "", "The display name for the Amazon Redshift IAM Identity Center application to change.")
		redshift_modifyRedshiftIdcApplicationCmd.Flags().String("identity-namespace", "", "The namespace for the Amazon Redshift IAM Identity Center application to change.")
		redshift_modifyRedshiftIdcApplicationCmd.Flags().String("redshift-idc-application-arn", "", "The ARN for the Redshift application that integrates with IAM Identity Center.")
		redshift_modifyRedshiftIdcApplicationCmd.Flags().String("service-integrations", "", "A collection of service integrations associated with the application.")
		redshift_modifyRedshiftIdcApplicationCmd.MarkFlagRequired("redshift-idc-application-arn")
	})
	redshiftCmd.AddCommand(redshift_modifyRedshiftIdcApplicationCmd)
}
