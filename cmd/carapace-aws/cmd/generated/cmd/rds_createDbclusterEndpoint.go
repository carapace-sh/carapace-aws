package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createDbclusterEndpointCmd = &cobra.Command{
	Use:   "create-dbcluster-endpoint",
	Short: "Creates a new custom endpoint and associates it with an Amazon Aurora DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createDbclusterEndpointCmd).Standalone()

	rds_createDbclusterEndpointCmd.Flags().String("dbcluster-endpoint-identifier", "", "The identifier to use for the new endpoint.")
	rds_createDbclusterEndpointCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier of the DB cluster associated with the endpoint.")
	rds_createDbclusterEndpointCmd.Flags().String("endpoint-type", "", "The type of the endpoint, one of: `READER`, `WRITER`, `ANY`.")
	rds_createDbclusterEndpointCmd.Flags().String("excluded-members", "", "List of DB instance identifiers that aren't part of the custom endpoint group.")
	rds_createDbclusterEndpointCmd.Flags().String("static-members", "", "List of DB instance identifiers that are part of the custom endpoint group.")
	rds_createDbclusterEndpointCmd.Flags().String("tags", "", "The tags to be assigned to the Amazon RDS resource.")
	rds_createDbclusterEndpointCmd.MarkFlagRequired("dbcluster-endpoint-identifier")
	rds_createDbclusterEndpointCmd.MarkFlagRequired("dbcluster-identifier")
	rds_createDbclusterEndpointCmd.MarkFlagRequired("endpoint-type")
	rdsCmd.AddCommand(rds_createDbclusterEndpointCmd)
}
