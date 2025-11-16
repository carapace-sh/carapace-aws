package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbproxyCmd = &cobra.Command{
	Use:   "modify-dbproxy",
	Short: "Changes the settings for an existing DB proxy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbproxyCmd).Standalone()

	rds_modifyDbproxyCmd.Flags().String("auth", "", "The new authentication settings for the `DBProxy`.")
	rds_modifyDbproxyCmd.Flags().String("dbproxy-name", "", "The identifier for the `DBProxy` to modify.")
	rds_modifyDbproxyCmd.Flags().String("debug-logging", "", "Specifies whether the proxy logs detailed connection and query information.")
	rds_modifyDbproxyCmd.Flags().String("default-auth-scheme", "", "The default authentication scheme that the proxy uses for client connections to the proxy and connections from the proxy to the underlying database.")
	rds_modifyDbproxyCmd.Flags().String("idle-client-timeout", "", "The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.")
	rds_modifyDbproxyCmd.Flags().String("new-dbproxy-name", "", "The new identifier for the `DBProxy`.")
	rds_modifyDbproxyCmd.Flags().String("require-tls", "", "Whether Transport Layer Security (TLS) encryption is required for connections to the proxy.")
	rds_modifyDbproxyCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in Amazon Web Services Secrets Manager.")
	rds_modifyDbproxyCmd.Flags().String("security-groups", "", "The new list of security groups for the `DBProxy`.")
	rds_modifyDbproxyCmd.MarkFlagRequired("dbproxy-name")
	rdsCmd.AddCommand(rds_modifyDbproxyCmd)
}
