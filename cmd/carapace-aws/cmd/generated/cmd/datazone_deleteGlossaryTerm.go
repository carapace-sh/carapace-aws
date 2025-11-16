package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteGlossaryTermCmd = &cobra.Command{
	Use:   "delete-glossary-term",
	Short: "Deletes a business glossary term in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteGlossaryTermCmd).Standalone()

	datazone_deleteGlossaryTermCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the business glossary term is deleted.")
	datazone_deleteGlossaryTermCmd.Flags().String("identifier", "", "The ID of the business glossary term that is deleted.")
	datazone_deleteGlossaryTermCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteGlossaryTermCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_deleteGlossaryTermCmd)
}
