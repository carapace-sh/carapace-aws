package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbproxyEndpointCmd = &cobra.Command{
	Use:   "modify-dbproxy-endpoint",
	Short: "Changes the settings for an existing DB proxy endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbproxyEndpointCmd).Standalone()

	rds_modifyDbproxyEndpointCmd.Flags().String("dbproxy-endpoint-name", "", "The name of the DB proxy sociated with the DB proxy endpoint that you want to modify.")
	rds_modifyDbproxyEndpointCmd.Flags().String("new-dbproxy-endpoint-name", "", "The new identifier for the `DBProxyEndpoint`.")
	rds_modifyDbproxyEndpointCmd.Flags().String("vpc-security-group-ids", "", "The VPC security group IDs for the DB proxy endpoint.")
	rds_modifyDbproxyEndpointCmd.MarkFlagRequired("dbproxy-endpoint-name")
	rdsCmd.AddCommand(rds_modifyDbproxyEndpointCmd)
}
