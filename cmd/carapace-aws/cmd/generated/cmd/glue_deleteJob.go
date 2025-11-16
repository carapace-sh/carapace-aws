package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteJobCmd = &cobra.Command{
	Use:   "delete-job",
	Short: "Deletes a specified job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteJobCmd).Standalone()

	glue_deleteJobCmd.Flags().String("job-name", "", "The name of the job definition to delete.")
	glue_deleteJobCmd.MarkFlagRequired("job-name")
	glueCmd.AddCommand(glue_deleteJobCmd)
}
