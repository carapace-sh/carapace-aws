package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getJobCmd = &cobra.Command{
	Use:   "get-job",
	Short: "Retrieves an existing job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getJobCmd).Standalone()

	glue_getJobCmd.Flags().String("job-name", "", "The name of the job definition to retrieve.")
	glue_getJobCmd.MarkFlagRequired("job-name")
	glueCmd.AddCommand(glue_getJobCmd)
}
