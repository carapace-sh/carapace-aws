package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_modifyDbclusterEndpointCmd = &cobra.Command{
	Use:   "modify-dbcluster-endpoint",
	Short: "Modifies the properties of an endpoint in an Amazon Neptune DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_modifyDbclusterEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_modifyDbclusterEndpointCmd).Standalone()

		neptune_modifyDbclusterEndpointCmd.Flags().String("dbcluster-endpoint-identifier", "", "The identifier of the endpoint to modify.")
		neptune_modifyDbclusterEndpointCmd.Flags().String("endpoint-type", "", "The type of the endpoint.")
		neptune_modifyDbclusterEndpointCmd.Flags().String("excluded-members", "", "List of DB instance identifiers that aren't part of the custom endpoint group.")
		neptune_modifyDbclusterEndpointCmd.Flags().String("static-members", "", "List of DB instance identifiers that are part of the custom endpoint group.")
		neptune_modifyDbclusterEndpointCmd.MarkFlagRequired("dbcluster-endpoint-identifier")
	})
	neptuneCmd.AddCommand(neptune_modifyDbclusterEndpointCmd)
}
