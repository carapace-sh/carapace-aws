package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateJobCmd = &cobra.Command{
	Use:   "update-job",
	Short: "Updates an existing job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateJobCmd).Standalone()

	glue_updateJobCmd.Flags().String("job-name", "", "The name of the job definition to update.")
	glue_updateJobCmd.Flags().String("job-update", "", "Specifies the values with which to update the job definition.")
	glue_updateJobCmd.MarkFlagRequired("job-name")
	glue_updateJobCmd.MarkFlagRequired("job-update")
	glueCmd.AddCommand(glue_updateJobCmd)
}
