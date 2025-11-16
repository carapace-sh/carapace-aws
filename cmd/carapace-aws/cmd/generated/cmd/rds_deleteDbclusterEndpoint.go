package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbclusterEndpointCmd = &cobra.Command{
	Use:   "delete-dbcluster-endpoint",
	Short: "Deletes a custom endpoint and removes it from an Amazon Aurora DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbclusterEndpointCmd).Standalone()

	rds_deleteDbclusterEndpointCmd.Flags().String("dbcluster-endpoint-identifier", "", "The identifier associated with the custom endpoint.")
	rds_deleteDbclusterEndpointCmd.MarkFlagRequired("dbcluster-endpoint-identifier")
	rdsCmd.AddCommand(rds_deleteDbclusterEndpointCmd)
}
