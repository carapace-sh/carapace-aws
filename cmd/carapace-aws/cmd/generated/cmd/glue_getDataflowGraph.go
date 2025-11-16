package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getDataflowGraphCmd = &cobra.Command{
	Use:   "get-dataflow-graph",
	Short: "Transforms a Python script into a directed acyclic graph (DAG).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getDataflowGraphCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getDataflowGraphCmd).Standalone()

		glue_getDataflowGraphCmd.Flags().String("python-script", "", "The Python script to transform.")
	})
	glueCmd.AddCommand(glue_getDataflowGraphCmd)
}
