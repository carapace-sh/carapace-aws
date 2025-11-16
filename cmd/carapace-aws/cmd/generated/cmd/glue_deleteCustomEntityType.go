package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteCustomEntityTypeCmd = &cobra.Command{
	Use:   "delete-custom-entity-type",
	Short: "Deletes a custom pattern by specifying its name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteCustomEntityTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteCustomEntityTypeCmd).Standalone()

		glue_deleteCustomEntityTypeCmd.Flags().String("name", "", "The name of the custom pattern that you want to delete.")
		glue_deleteCustomEntityTypeCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_deleteCustomEntityTypeCmd)
}
