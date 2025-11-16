package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_importNotebookCmd = &cobra.Command{
	Use:   "import-notebook",
	Short: "Imports a single `ipynb` file to a Spark enabled workgroup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_importNotebookCmd).Standalone()

	athena_importNotebookCmd.Flags().String("client-request-token", "", "A unique case-sensitive string used to ensure the request to import the notebook is idempotent (executes only once).")
	athena_importNotebookCmd.Flags().String("name", "", "The name of the notebook to import.")
	athena_importNotebookCmd.Flags().String("notebook-s3-location-uri", "", "A URI that specifies the Amazon S3 location of a notebook file in `ipynb` format.")
	athena_importNotebookCmd.Flags().String("payload", "", "The notebook content to be imported.")
	athena_importNotebookCmd.Flags().String("type", "", "The notebook content type.")
	athena_importNotebookCmd.Flags().String("work-group", "", "The name of the Spark enabled workgroup to import the notebook to.")
	athena_importNotebookCmd.MarkFlagRequired("name")
	athena_importNotebookCmd.MarkFlagRequired("type")
	athena_importNotebookCmd.MarkFlagRequired("work-group")
	athenaCmd.AddCommand(athena_importNotebookCmd)
}
