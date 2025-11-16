package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_deregisterJobDefinitionCmd = &cobra.Command{
	Use:   "deregister-job-definition",
	Short: "Deregisters an Batch job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_deregisterJobDefinitionCmd).Standalone()

	batch_deregisterJobDefinitionCmd.Flags().String("job-definition", "", "The name and revision (`name:revision`) or full Amazon Resource Name (ARN) of the job definition to deregister.")
	batch_deregisterJobDefinitionCmd.MarkFlagRequired("job-definition")
	batchCmd.AddCommand(batch_deregisterJobDefinitionCmd)
}
