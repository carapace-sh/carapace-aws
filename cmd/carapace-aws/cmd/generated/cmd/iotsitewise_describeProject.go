package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeProjectCmd = &cobra.Command{
	Use:   "describe-project",
	Short: "Retrieves information about a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeProjectCmd).Standalone()

	iotsitewise_describeProjectCmd.Flags().String("project-id", "", "The ID of the project.")
	iotsitewise_describeProjectCmd.MarkFlagRequired("project-id")
	iotsitewiseCmd.AddCommand(iotsitewise_describeProjectCmd)
}
