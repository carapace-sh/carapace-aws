package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_getLaunchCmd = &cobra.Command{
	Use:   "get-launch",
	Short: "Returns the details about one launch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_getLaunchCmd).Standalone()

	evidently_getLaunchCmd.Flags().String("launch", "", "The name of the launch that you want to see the details of.")
	evidently_getLaunchCmd.Flags().String("project", "", "The name or ARN of the project that contains the launch.")
	evidently_getLaunchCmd.MarkFlagRequired("launch")
	evidently_getLaunchCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_getLaunchCmd)
}
