package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_deleteProjectCmd = &cobra.Command{
	Use:   "delete-project",
	Short: "Deletes an AWS Device Farm project, given the project ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_deleteProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_deleteProjectCmd).Standalone()

		devicefarm_deleteProjectCmd.Flags().String("arn", "", "Represents the Amazon Resource Name (ARN) of the Device Farm project to delete.")
		devicefarm_deleteProjectCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_deleteProjectCmd)
}
