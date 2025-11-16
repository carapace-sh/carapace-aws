package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopAutoMljobCmd = &cobra.Command{
	Use:   "stop-auto-mljob",
	Short: "A method for forcing a running job to shut down.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopAutoMljobCmd).Standalone()

	sagemaker_stopAutoMljobCmd.Flags().String("auto-mljob-name", "", "The name of the object you are requesting.")
	sagemaker_stopAutoMljobCmd.MarkFlagRequired("auto-mljob-name")
	sagemakerCmd.AddCommand(sagemaker_stopAutoMljobCmd)
}
