package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createCustomEntityTypeCmd = &cobra.Command{
	Use:   "create-custom-entity-type",
	Short: "Creates a custom pattern that is used to detect sensitive data across the columns and rows of your structured data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createCustomEntityTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createCustomEntityTypeCmd).Standalone()

		glue_createCustomEntityTypeCmd.Flags().String("context-words", "", "A list of context words.")
		glue_createCustomEntityTypeCmd.Flags().String("name", "", "A name for the custom pattern that allows it to be retrieved or deleted later.")
		glue_createCustomEntityTypeCmd.Flags().String("regex-string", "", "A regular expression string that is used for detecting sensitive data in a custom pattern.")
		glue_createCustomEntityTypeCmd.Flags().String("tags", "", "A list of tags applied to the custom entity type.")
		glue_createCustomEntityTypeCmd.MarkFlagRequired("name")
		glue_createCustomEntityTypeCmd.MarkFlagRequired("regex-string")
	})
	glueCmd.AddCommand(glue_createCustomEntityTypeCmd)
}
