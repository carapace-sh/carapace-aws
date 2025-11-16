package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteEndpointAccessCmd = &cobra.Command{
	Use:   "delete-endpoint-access",
	Short: "Deletes a Redshift-managed VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteEndpointAccessCmd).Standalone()

	redshift_deleteEndpointAccessCmd.Flags().String("endpoint-name", "", "The Redshift-managed VPC endpoint to delete.")
	redshift_deleteEndpointAccessCmd.MarkFlagRequired("endpoint-name")
	redshiftCmd.AddCommand(redshift_deleteEndpointAccessCmd)
}
