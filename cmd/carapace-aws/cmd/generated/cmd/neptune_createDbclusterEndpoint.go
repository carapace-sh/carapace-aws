package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_createDbclusterEndpointCmd = &cobra.Command{
	Use:   "create-dbcluster-endpoint",
	Short: "Creates a new custom endpoint and associates it with an Amazon Neptune DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_createDbclusterEndpointCmd).Standalone()

	neptune_createDbclusterEndpointCmd.Flags().String("dbcluster-endpoint-identifier", "", "The identifier to use for the new endpoint.")
	neptune_createDbclusterEndpointCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier of the DB cluster associated with the endpoint.")
	neptune_createDbclusterEndpointCmd.Flags().String("endpoint-type", "", "The type of the endpoint.")
	neptune_createDbclusterEndpointCmd.Flags().String("excluded-members", "", "List of DB instance identifiers that aren't part of the custom endpoint group.")
	neptune_createDbclusterEndpointCmd.Flags().String("static-members", "", "List of DB instance identifiers that are part of the custom endpoint group.")
	neptune_createDbclusterEndpointCmd.Flags().String("tags", "", "The tags to be assigned to the Amazon Neptune resource.")
	neptune_createDbclusterEndpointCmd.MarkFlagRequired("dbcluster-endpoint-identifier")
	neptune_createDbclusterEndpointCmd.MarkFlagRequired("dbcluster-identifier")
	neptune_createDbclusterEndpointCmd.MarkFlagRequired("endpoint-type")
	neptuneCmd.AddCommand(neptune_createDbclusterEndpointCmd)
}
