package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getWorkGroupCmd = &cobra.Command{
	Use:   "get-work-group",
	Short: "Returns information about the workgroup with the specified name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getWorkGroupCmd).Standalone()

	athena_getWorkGroupCmd.Flags().String("work-group", "", "The name of the workgroup.")
	athena_getWorkGroupCmd.MarkFlagRequired("work-group")
	athenaCmd.AddCommand(athena_getWorkGroupCmd)
}
