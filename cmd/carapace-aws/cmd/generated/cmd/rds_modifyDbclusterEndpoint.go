package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbclusterEndpointCmd = &cobra.Command{
	Use:   "modify-dbcluster-endpoint",
	Short: "Modifies the properties of an endpoint in an Amazon Aurora DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbclusterEndpointCmd).Standalone()

	rds_modifyDbclusterEndpointCmd.Flags().String("dbcluster-endpoint-identifier", "", "The identifier of the endpoint to modify.")
	rds_modifyDbclusterEndpointCmd.Flags().String("endpoint-type", "", "The type of the endpoint.")
	rds_modifyDbclusterEndpointCmd.Flags().String("excluded-members", "", "List of DB instance identifiers that aren't part of the custom endpoint group.")
	rds_modifyDbclusterEndpointCmd.Flags().String("static-members", "", "List of DB instance identifiers that are part of the custom endpoint group.")
	rds_modifyDbclusterEndpointCmd.MarkFlagRequired("dbcluster-endpoint-identifier")
	rdsCmd.AddCommand(rds_modifyDbclusterEndpointCmd)
}
