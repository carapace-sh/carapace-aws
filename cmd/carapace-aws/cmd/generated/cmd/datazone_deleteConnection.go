package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "Deletes and connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteConnectionCmd).Standalone()

	datazone_deleteConnectionCmd.Flags().String("domain-identifier", "", "The ID of the domain where the connection is deleted.")
	datazone_deleteConnectionCmd.Flags().String("identifier", "", "The ID of the connection that is deleted.")
	datazone_deleteConnectionCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteConnectionCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_deleteConnectionCmd)
}
