package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getDevEndpointCmd = &cobra.Command{
	Use:   "get-dev-endpoint",
	Short: "Retrieves information about a specified development endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getDevEndpointCmd).Standalone()

	glue_getDevEndpointCmd.Flags().String("endpoint-name", "", "Name of the `DevEndpoint` to retrieve information for.")
	glue_getDevEndpointCmd.MarkFlagRequired("endpoint-name")
	glueCmd.AddCommand(glue_getDevEndpointCmd)
}
