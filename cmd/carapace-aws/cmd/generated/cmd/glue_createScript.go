package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createScriptCmd = &cobra.Command{
	Use:   "create-script",
	Short: "Transforms a directed acyclic graph (DAG) into code.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createScriptCmd).Standalone()

	glue_createScriptCmd.Flags().String("dag-edges", "", "A list of the edges in the DAG.")
	glue_createScriptCmd.Flags().String("dag-nodes", "", "A list of the nodes in the DAG.")
	glue_createScriptCmd.Flags().String("language", "", "The programming language of the resulting code from the DAG.")
	glueCmd.AddCommand(glue_createScriptCmd)
}
