package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateDevEndpointCmd = &cobra.Command{
	Use:   "update-dev-endpoint",
	Short: "Updates a specified development endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateDevEndpointCmd).Standalone()

	glue_updateDevEndpointCmd.Flags().String("add-arguments", "", "The map of arguments to add the map of arguments used to configure the `DevEndpoint`.")
	glue_updateDevEndpointCmd.Flags().String("add-public-keys", "", "The list of public keys for the `DevEndpoint` to use.")
	glue_updateDevEndpointCmd.Flags().String("custom-libraries", "", "Custom Python or Java libraries to be loaded in the `DevEndpoint`.")
	glue_updateDevEndpointCmd.Flags().String("delete-arguments", "", "The list of argument keys to be deleted from the map of arguments used to configure the `DevEndpoint`.")
	glue_updateDevEndpointCmd.Flags().String("delete-public-keys", "", "The list of public keys to be deleted from the `DevEndpoint`.")
	glue_updateDevEndpointCmd.Flags().String("endpoint-name", "", "The name of the `DevEndpoint` to be updated.")
	glue_updateDevEndpointCmd.Flags().String("public-key", "", "The public key for the `DevEndpoint` to use.")
	glue_updateDevEndpointCmd.Flags().String("update-etl-libraries", "", "`True` if the list of custom libraries to be loaded in the development endpoint needs to be updated, or `False` if otherwise.")
	glue_updateDevEndpointCmd.MarkFlagRequired("endpoint-name")
	glueCmd.AddCommand(glue_updateDevEndpointCmd)
}
