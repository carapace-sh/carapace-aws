package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createDbproxyCmd = &cobra.Command{
	Use:   "create-dbproxy",
	Short: "Creates a new DB proxy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createDbproxyCmd).Standalone()

	rds_createDbproxyCmd.Flags().String("auth", "", "The authorization mechanism that the proxy uses.")
	rds_createDbproxyCmd.Flags().String("dbproxy-name", "", "The identifier for the proxy.")
	rds_createDbproxyCmd.Flags().Bool("debug-logging", false, "Specifies whether the proxy logs detailed connection and query information.")
	rds_createDbproxyCmd.Flags().String("default-auth-scheme", "", "The default authentication scheme that the proxy uses for client connections to the proxy and connections from the proxy to the underlying database.")
	rds_createDbproxyCmd.Flags().String("endpoint-network-type", "", "The network type of the DB proxy endpoint.")
	rds_createDbproxyCmd.Flags().String("engine-family", "", "The kinds of databases that the proxy can connect to.")
	rds_createDbproxyCmd.Flags().String("idle-client-timeout", "", "The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.")
	rds_createDbproxyCmd.Flags().Bool("no-debug-logging", false, "Specifies whether the proxy logs detailed connection and query information.")
	rds_createDbproxyCmd.Flags().Bool("no-require-tls", false, "Specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.")
	rds_createDbproxyCmd.Flags().Bool("require-tls", false, "Specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.")
	rds_createDbproxyCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in Amazon Web Services Secrets Manager.")
	rds_createDbproxyCmd.Flags().String("tags", "", "An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.")
	rds_createDbproxyCmd.Flags().String("target-connection-network-type", "", "The network type that the proxy uses to connect to the target database.")
	rds_createDbproxyCmd.Flags().String("vpc-security-group-ids", "", "One or more VPC security group IDs to associate with the new proxy.")
	rds_createDbproxyCmd.Flags().String("vpc-subnet-ids", "", "One or more VPC subnet IDs to associate with the new proxy.")
	rds_createDbproxyCmd.MarkFlagRequired("dbproxy-name")
	rds_createDbproxyCmd.MarkFlagRequired("engine-family")
	rds_createDbproxyCmd.Flag("no-debug-logging").Hidden = true
	rds_createDbproxyCmd.Flag("no-require-tls").Hidden = true
	rds_createDbproxyCmd.MarkFlagRequired("role-arn")
	rds_createDbproxyCmd.MarkFlagRequired("vpc-subnet-ids")
	rdsCmd.AddCommand(rds_createDbproxyCmd)
}
