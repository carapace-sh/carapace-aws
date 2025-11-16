package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createGlossaryTermCmd = &cobra.Command{
	Use:   "create-glossary-term",
	Short: "Creates a business glossary term.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createGlossaryTermCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createGlossaryTermCmd).Standalone()

		datazone_createGlossaryTermCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createGlossaryTermCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which this business glossary term is created.")
		datazone_createGlossaryTermCmd.Flags().String("glossary-identifier", "", "The ID of the business glossary in which this term is created.")
		datazone_createGlossaryTermCmd.Flags().String("long-description", "", "The long description of this business glossary term.")
		datazone_createGlossaryTermCmd.Flags().String("name", "", "The name of this business glossary term.")
		datazone_createGlossaryTermCmd.Flags().String("short-description", "", "The short description of this business glossary term.")
		datazone_createGlossaryTermCmd.Flags().String("status", "", "The status of this business glossary term.")
		datazone_createGlossaryTermCmd.Flags().String("term-relations", "", "The term relations of this business glossary term.")
		datazone_createGlossaryTermCmd.MarkFlagRequired("domain-identifier")
		datazone_createGlossaryTermCmd.MarkFlagRequired("glossary-identifier")
		datazone_createGlossaryTermCmd.MarkFlagRequired("name")
	})
	datazoneCmd.AddCommand(datazone_createGlossaryTermCmd)
}
