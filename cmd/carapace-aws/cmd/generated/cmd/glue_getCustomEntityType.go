package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getCustomEntityTypeCmd = &cobra.Command{
	Use:   "get-custom-entity-type",
	Short: "Retrieves the details of a custom pattern by specifying its name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getCustomEntityTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getCustomEntityTypeCmd).Standalone()

		glue_getCustomEntityTypeCmd.Flags().String("name", "", "The name of the custom pattern that you want to retrieve.")
		glue_getCustomEntityTypeCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_getCustomEntityTypeCmd)
}
