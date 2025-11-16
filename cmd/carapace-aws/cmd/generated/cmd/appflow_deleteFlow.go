package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_deleteFlowCmd = &cobra.Command{
	Use:   "delete-flow",
	Short: "Enables your application to delete an existing flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_deleteFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_deleteFlowCmd).Standalone()

		appflow_deleteFlowCmd.Flags().String("flow-name", "", "The specified name of the flow.")
		appflow_deleteFlowCmd.Flags().Bool("force-delete", false, "Indicates whether Amazon AppFlow should delete the flow, even if it is currently in use.")
		appflow_deleteFlowCmd.Flags().Bool("no-force-delete", false, "Indicates whether Amazon AppFlow should delete the flow, even if it is currently in use.")
		appflow_deleteFlowCmd.MarkFlagRequired("flow-name")
		appflow_deleteFlowCmd.Flag("no-force-delete").Hidden = true
	})
	appflowCmd.AddCommand(appflow_deleteFlowCmd)
}
