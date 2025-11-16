package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createDbproxyEndpointCmd = &cobra.Command{
	Use:   "create-dbproxy-endpoint",
	Short: "Creates a `DBProxyEndpoint`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createDbproxyEndpointCmd).Standalone()

	rds_createDbproxyEndpointCmd.Flags().String("dbproxy-endpoint-name", "", "The name of the DB proxy endpoint to create.")
	rds_createDbproxyEndpointCmd.Flags().String("dbproxy-name", "", "The name of the DB proxy associated with the DB proxy endpoint that you create.")
	rds_createDbproxyEndpointCmd.Flags().String("endpoint-network-type", "", "The network type of the DB proxy endpoint.")
	rds_createDbproxyEndpointCmd.Flags().String("tags", "", "")
	rds_createDbproxyEndpointCmd.Flags().String("target-role", "", "The role of the DB proxy endpoint.")
	rds_createDbproxyEndpointCmd.Flags().String("vpc-security-group-ids", "", "The VPC security group IDs for the DB proxy endpoint that you create.")
	rds_createDbproxyEndpointCmd.Flags().String("vpc-subnet-ids", "", "The VPC subnet IDs for the DB proxy endpoint that you create.")
	rds_createDbproxyEndpointCmd.MarkFlagRequired("dbproxy-endpoint-name")
	rds_createDbproxyEndpointCmd.MarkFlagRequired("dbproxy-name")
	rds_createDbproxyEndpointCmd.MarkFlagRequired("vpc-subnet-ids")
	rdsCmd.AddCommand(rds_createDbproxyEndpointCmd)
}
