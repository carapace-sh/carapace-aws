package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getMltaskRunCmd = &cobra.Command{
	Use:   "get-mltask-run",
	Short: "Gets details for a specific task run on a machine learning transform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getMltaskRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getMltaskRunCmd).Standalone()

		glue_getMltaskRunCmd.Flags().String("task-run-id", "", "The unique identifier of the task run.")
		glue_getMltaskRunCmd.Flags().String("transform-id", "", "The unique identifier of the machine learning transform.")
		glue_getMltaskRunCmd.MarkFlagRequired("task-run-id")
		glue_getMltaskRunCmd.MarkFlagRequired("transform-id")
	})
	glueCmd.AddCommand(glue_getMltaskRunCmd)
}
