package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_createNotebookCmd = &cobra.Command{
	Use:   "create-notebook",
	Short: "Creates an empty `ipynb` file in the specified Apache Spark enabled workgroup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_createNotebookCmd).Standalone()

	athena_createNotebookCmd.Flags().String("client-request-token", "", "A unique case-sensitive string used to ensure the request to create the notebook is idempotent (executes only once).")
	athena_createNotebookCmd.Flags().String("name", "", "The name of the `ipynb` file to be created in the Spark workgroup, without the `.ipynb` extension.")
	athena_createNotebookCmd.Flags().String("work-group", "", "The name of the Spark enabled workgroup in which the notebook will be created.")
	athena_createNotebookCmd.MarkFlagRequired("name")
	athena_createNotebookCmd.MarkFlagRequired("work-group")
	athenaCmd.AddCommand(athena_createNotebookCmd)
}
