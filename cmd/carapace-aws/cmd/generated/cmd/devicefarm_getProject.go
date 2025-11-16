package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getProjectCmd = &cobra.Command{
	Use:   "get-project",
	Short: "Gets information about a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getProjectCmd).Standalone()

	devicefarm_getProjectCmd.Flags().String("arn", "", "The project's ARN.")
	devicefarm_getProjectCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_getProjectCmd)
}
