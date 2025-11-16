package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_deleteDbclusterEndpointCmd = &cobra.Command{
	Use:   "delete-dbcluster-endpoint",
	Short: "Deletes a custom endpoint and removes it from an Amazon Neptune DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_deleteDbclusterEndpointCmd).Standalone()

	neptune_deleteDbclusterEndpointCmd.Flags().String("dbcluster-endpoint-identifier", "", "The identifier associated with the custom endpoint.")
	neptune_deleteDbclusterEndpointCmd.MarkFlagRequired("dbcluster-endpoint-identifier")
	neptuneCmd.AddCommand(neptune_deleteDbclusterEndpointCmd)
}
