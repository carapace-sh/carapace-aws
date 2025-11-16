package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_createNamespaceCmd = &cobra.Command{
	Use:   "create-namespace",
	Short: "Creates a namespace in Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_createNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_createNamespaceCmd).Standalone()

		redshiftServerless_createNamespaceCmd.Flags().String("admin-password-secret-kms-key-id", "", "The ID of the Key Management Service (KMS) key used to encrypt and store the namespace's admin credentials secret.")
		redshiftServerless_createNamespaceCmd.Flags().String("admin-user-password", "", "The password of the administrator for the first database created in the namespace.")
		redshiftServerless_createNamespaceCmd.Flags().String("admin-username", "", "The username of the administrator for the first database created in the namespace.")
		redshiftServerless_createNamespaceCmd.Flags().String("db-name", "", "The name of the first database created in the namespace.")
		redshiftServerless_createNamespaceCmd.Flags().String("default-iam-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace.")
		redshiftServerless_createNamespaceCmd.Flags().String("iam-roles", "", "A list of IAM roles to associate with the namespace.")
		redshiftServerless_createNamespaceCmd.Flags().String("kms-key-id", "", "The ID of the Amazon Web Services Key Management Service key used to encrypt your data.")
		redshiftServerless_createNamespaceCmd.Flags().String("log-exports", "", "The types of logs the namespace can export.")
		redshiftServerless_createNamespaceCmd.Flags().Bool("manage-admin-password", false, "If `true`, Amazon Redshift uses Secrets Manager to manage the namespace's admin credentials.")
		redshiftServerless_createNamespaceCmd.Flags().String("namespace-name", "", "The name of the namespace.")
		redshiftServerless_createNamespaceCmd.Flags().Bool("no-manage-admin-password", false, "If `true`, Amazon Redshift uses Secrets Manager to manage the namespace's admin credentials.")
		redshiftServerless_createNamespaceCmd.Flags().String("redshift-idc-application-arn", "", "The ARN for the Redshift application that integrates with IAM Identity Center.")
		redshiftServerless_createNamespaceCmd.Flags().String("tags", "", "A list of tag instances.")
		redshiftServerless_createNamespaceCmd.MarkFlagRequired("namespace-name")
		redshiftServerless_createNamespaceCmd.Flag("no-manage-admin-password").Hidden = true
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_createNamespaceCmd)
}
