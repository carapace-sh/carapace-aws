package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_deleteWorkGroupCmd = &cobra.Command{
	Use:   "delete-work-group",
	Short: "Deletes the workgroup with the specified name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_deleteWorkGroupCmd).Standalone()

	athena_deleteWorkGroupCmd.Flags().String("recursive-delete-option", "", "The option to delete the workgroup and its contents even if the workgroup contains any named queries, query executions, or notebooks.")
	athena_deleteWorkGroupCmd.Flags().String("work-group", "", "The unique name of the workgroup to delete.")
	athena_deleteWorkGroupCmd.MarkFlagRequired("work-group")
	athenaCmd.AddCommand(athena_deleteWorkGroupCmd)
}
