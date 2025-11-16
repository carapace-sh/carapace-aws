package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createGlossaryCmd = &cobra.Command{
	Use:   "create-glossary",
	Short: "Creates an Amazon DataZone business glossary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createGlossaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createGlossaryCmd).Standalone()

		datazone_createGlossaryCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createGlossaryCmd.Flags().String("description", "", "The description of this business glossary.")
		datazone_createGlossaryCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which this business glossary is created.")
		datazone_createGlossaryCmd.Flags().String("name", "", "The name of this business glossary.")
		datazone_createGlossaryCmd.Flags().String("owning-project-identifier", "", "The ID of the project that currently owns business glossary.")
		datazone_createGlossaryCmd.Flags().String("status", "", "The status of this business glossary.")
		datazone_createGlossaryCmd.Flags().String("usage-restrictions", "", "The usage restriction of the restricted glossary.")
		datazone_createGlossaryCmd.MarkFlagRequired("domain-identifier")
		datazone_createGlossaryCmd.MarkFlagRequired("name")
		datazone_createGlossaryCmd.MarkFlagRequired("owning-project-identifier")
	})
	datazoneCmd.AddCommand(datazone_createGlossaryCmd)
}
