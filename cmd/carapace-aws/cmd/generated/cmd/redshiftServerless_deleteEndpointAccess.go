package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_deleteEndpointAccessCmd = &cobra.Command{
	Use:   "delete-endpoint-access",
	Short: "Deletes an Amazon Redshift Serverless managed VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_deleteEndpointAccessCmd).Standalone()

	redshiftServerless_deleteEndpointAccessCmd.Flags().String("endpoint-name", "", "The name of the VPC endpoint to delete.")
	redshiftServerless_deleteEndpointAccessCmd.MarkFlagRequired("endpoint-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_deleteEndpointAccessCmd)
}
