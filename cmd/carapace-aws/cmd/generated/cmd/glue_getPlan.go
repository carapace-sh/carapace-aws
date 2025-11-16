package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getPlanCmd = &cobra.Command{
	Use:   "get-plan",
	Short: "Gets code to perform a specified mapping.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getPlanCmd).Standalone()

	glue_getPlanCmd.Flags().String("additional-plan-options-map", "", "A map to hold additional optional key-value parameters.")
	glue_getPlanCmd.Flags().String("language", "", "The programming language of the code to perform the mapping.")
	glue_getPlanCmd.Flags().String("location", "", "The parameters for the mapping.")
	glue_getPlanCmd.Flags().String("mapping", "", "The list of mappings from a source table to target tables.")
	glue_getPlanCmd.Flags().String("sinks", "", "The target tables.")
	glue_getPlanCmd.Flags().String("source", "", "The source table.")
	glue_getPlanCmd.MarkFlagRequired("mapping")
	glue_getPlanCmd.MarkFlagRequired("source")
	glueCmd.AddCommand(glue_getPlanCmd)
}
