package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateConnectionCmd = &cobra.Command{
	Use:   "update-connection",
	Short: "Updates a connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateConnectionCmd).Standalone()

	datazone_updateConnectionCmd.Flags().String("aws-location", "", "The location where a connection is to be updated.")
	datazone_updateConnectionCmd.Flags().String("description", "", "The description of a connection.")
	datazone_updateConnectionCmd.Flags().String("domain-identifier", "", "The ID of the domain where a connection is to be updated.")
	datazone_updateConnectionCmd.Flags().String("identifier", "", "The ID of the connection to be updated.")
	datazone_updateConnectionCmd.Flags().String("props", "", "The connection props.")
	datazone_updateConnectionCmd.MarkFlagRequired("domain-identifier")
	datazone_updateConnectionCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_updateConnectionCmd)
}
