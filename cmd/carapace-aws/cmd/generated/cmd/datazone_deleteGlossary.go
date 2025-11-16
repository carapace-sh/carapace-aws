package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteGlossaryCmd = &cobra.Command{
	Use:   "delete-glossary",
	Short: "Deletes a business glossary in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteGlossaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteGlossaryCmd).Standalone()

		datazone_deleteGlossaryCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the business glossary is deleted.")
		datazone_deleteGlossaryCmd.Flags().String("identifier", "", "The ID of the business glossary that is deleted.")
		datazone_deleteGlossaryCmd.MarkFlagRequired("domain-identifier")
		datazone_deleteGlossaryCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_deleteGlossaryCmd)
}
