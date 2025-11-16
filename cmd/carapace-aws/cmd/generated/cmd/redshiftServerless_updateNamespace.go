package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_updateNamespaceCmd = &cobra.Command{
	Use:   "update-namespace",
	Short: "Updates a namespace with the specified settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_updateNamespaceCmd).Standalone()

	redshiftServerless_updateNamespaceCmd.Flags().String("admin-password-secret-kms-key-id", "", "The ID of the Key Management Service (KMS) key used to encrypt and store the namespace's admin credentials secret.")
	redshiftServerless_updateNamespaceCmd.Flags().String("admin-user-password", "", "The password of the administrator for the first database created in the namespace.")
	redshiftServerless_updateNamespaceCmd.Flags().String("admin-username", "", "The username of the administrator for the first database created in the namespace.")
	redshiftServerless_updateNamespaceCmd.Flags().String("default-iam-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace.")
	redshiftServerless_updateNamespaceCmd.Flags().String("iam-roles", "", "A list of IAM roles to associate with the namespace.")
	redshiftServerless_updateNamespaceCmd.Flags().String("kms-key-id", "", "The ID of the Amazon Web Services Key Management Service key used to encrypt your data.")
	redshiftServerless_updateNamespaceCmd.Flags().String("log-exports", "", "The types of logs the namespace can export.")
	redshiftServerless_updateNamespaceCmd.Flags().Bool("manage-admin-password", false, "If `true`, Amazon Redshift uses Secrets Manager to manage the namespace's admin credentials.")
	redshiftServerless_updateNamespaceCmd.Flags().String("namespace-name", "", "The name of the namespace to update.")
	redshiftServerless_updateNamespaceCmd.Flags().Bool("no-manage-admin-password", false, "If `true`, Amazon Redshift uses Secrets Manager to manage the namespace's admin credentials.")
	redshiftServerless_updateNamespaceCmd.MarkFlagRequired("namespace-name")
	redshiftServerless_updateNamespaceCmd.Flag("no-manage-admin-password").Hidden = true
	redshiftServerlessCmd.AddCommand(redshiftServerless_updateNamespaceCmd)
}
