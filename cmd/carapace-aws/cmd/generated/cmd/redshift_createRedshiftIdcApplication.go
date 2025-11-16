package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createRedshiftIdcApplicationCmd = &cobra.Command{
	Use:   "create-redshift-idc-application",
	Short: "Creates an Amazon Redshift application for use with IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createRedshiftIdcApplicationCmd).Standalone()

	redshift_createRedshiftIdcApplicationCmd.Flags().String("authorized-token-issuer-list", "", "The token issuer list for the Amazon Redshift IAM Identity Center application instance.")
	redshift_createRedshiftIdcApplicationCmd.Flags().String("iam-role-arn", "", "The IAM role ARN for the Amazon Redshift IAM Identity Center application instance.")
	redshift_createRedshiftIdcApplicationCmd.Flags().String("idc-display-name", "", "The display name for the Amazon Redshift IAM Identity Center application instance.")
	redshift_createRedshiftIdcApplicationCmd.Flags().String("idc-instance-arn", "", "The Amazon resource name (ARN) of the IAM Identity Center instance where Amazon Redshift creates a new managed application.")
	redshift_createRedshiftIdcApplicationCmd.Flags().String("identity-namespace", "", "The namespace for the Amazon Redshift IAM Identity Center application instance.")
	redshift_createRedshiftIdcApplicationCmd.Flags().String("redshift-idc-application-name", "", "The name of the Redshift application in IAM Identity Center.")
	redshift_createRedshiftIdcApplicationCmd.Flags().String("service-integrations", "", "A collection of service integrations for the Redshift IAM Identity Center application.")
	redshift_createRedshiftIdcApplicationCmd.Flags().String("sso-tag-keys", "", "A list of tags keys that Redshift Identity Center applications copy to IAM Identity Center.")
	redshift_createRedshiftIdcApplicationCmd.Flags().String("tags", "", "A list of tags.")
	redshift_createRedshiftIdcApplicationCmd.MarkFlagRequired("iam-role-arn")
	redshift_createRedshiftIdcApplicationCmd.MarkFlagRequired("idc-display-name")
	redshift_createRedshiftIdcApplicationCmd.MarkFlagRequired("idc-instance-arn")
	redshift_createRedshiftIdcApplicationCmd.MarkFlagRequired("redshift-idc-application-name")
	redshiftCmd.AddCommand(redshift_createRedshiftIdcApplicationCmd)
}
