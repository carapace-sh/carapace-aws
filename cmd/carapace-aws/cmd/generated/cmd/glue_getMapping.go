package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getMappingCmd = &cobra.Command{
	Use:   "get-mapping",
	Short: "Creates mappings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getMappingCmd).Standalone()

		glue_getMappingCmd.Flags().String("location", "", "Parameters for the mapping.")
		glue_getMappingCmd.Flags().String("sinks", "", "A list of target tables.")
		glue_getMappingCmd.Flags().String("source", "", "Specifies the source table.")
		glue_getMappingCmd.MarkFlagRequired("source")
	})
	glueCmd.AddCommand(glue_getMappingCmd)
}
