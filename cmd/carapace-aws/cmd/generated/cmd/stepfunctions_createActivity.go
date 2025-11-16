package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_createActivityCmd = &cobra.Command{
	Use:   "create-activity",
	Short: "Creates an activity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_createActivityCmd).Standalone()

	stepfunctions_createActivityCmd.Flags().String("encryption-configuration", "", "Settings to configure server-side encryption.")
	stepfunctions_createActivityCmd.Flags().String("name", "", "The name of the activity to create.")
	stepfunctions_createActivityCmd.Flags().String("tags", "", "The list of tags to add to a resource.")
	stepfunctions_createActivityCmd.MarkFlagRequired("name")
	stepfunctionsCmd.AddCommand(stepfunctions_createActivityCmd)
}
