package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deleteParameterCmd = &cobra.Command{
	Use:   "delete-parameter",
	Short: "Delete a parameter from the system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deleteParameterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_deleteParameterCmd).Standalone()

		ssm_deleteParameterCmd.Flags().String("name", "", "The name of the parameter to delete.")
		ssm_deleteParameterCmd.MarkFlagRequired("name")
	})
	ssmCmd.AddCommand(ssm_deleteParameterCmd)
}
