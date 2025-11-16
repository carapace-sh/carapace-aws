package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_createWorkGroupCmd = &cobra.Command{
	Use:   "create-work-group",
	Short: "Creates a workgroup with the specified name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_createWorkGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_createWorkGroupCmd).Standalone()

		athena_createWorkGroupCmd.Flags().String("configuration", "", "Contains configuration information for creating an Athena SQL workgroup or Spark enabled Athena workgroup.")
		athena_createWorkGroupCmd.Flags().String("description", "", "The workgroup description.")
		athena_createWorkGroupCmd.Flags().String("name", "", "The workgroup name.")
		athena_createWorkGroupCmd.Flags().String("tags", "", "A list of comma separated tags to add to the workgroup that is created.")
		athena_createWorkGroupCmd.MarkFlagRequired("name")
	})
	athenaCmd.AddCommand(athena_createWorkGroupCmd)
}
