package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbproxyEndpointCmd = &cobra.Command{
	Use:   "delete-dbproxy-endpoint",
	Short: "Deletes a `DBProxyEndpoint`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbproxyEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteDbproxyEndpointCmd).Standalone()

		rds_deleteDbproxyEndpointCmd.Flags().String("dbproxy-endpoint-name", "", "The name of the DB proxy endpoint to delete.")
		rds_deleteDbproxyEndpointCmd.MarkFlagRequired("dbproxy-endpoint-name")
	})
	rdsCmd.AddCommand(rds_deleteDbproxyEndpointCmd)
}
