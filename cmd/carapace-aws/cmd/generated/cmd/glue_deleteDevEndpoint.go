package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteDevEndpointCmd = &cobra.Command{
	Use:   "delete-dev-endpoint",
	Short: "Deletes a specified development endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteDevEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteDevEndpointCmd).Standalone()

		glue_deleteDevEndpointCmd.Flags().String("endpoint-name", "", "The name of the `DevEndpoint`.")
		glue_deleteDevEndpointCmd.MarkFlagRequired("endpoint-name")
	})
	glueCmd.AddCommand(glue_deleteDevEndpointCmd)
}
