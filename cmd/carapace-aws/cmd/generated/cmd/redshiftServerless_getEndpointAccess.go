package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getEndpointAccessCmd = &cobra.Command{
	Use:   "get-endpoint-access",
	Short: "Returns information, such as the name, about a VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getEndpointAccessCmd).Standalone()

	redshiftServerless_getEndpointAccessCmd.Flags().String("endpoint-name", "", "The name of the VPC endpoint to return information for.")
	redshiftServerless_getEndpointAccessCmd.MarkFlagRequired("endpoint-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_getEndpointAccessCmd)
}
