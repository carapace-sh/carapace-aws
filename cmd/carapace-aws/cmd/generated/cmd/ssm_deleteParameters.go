package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deleteParametersCmd = &cobra.Command{
	Use:   "delete-parameters",
	Short: "Delete a list of parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deleteParametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_deleteParametersCmd).Standalone()

		ssm_deleteParametersCmd.Flags().String("names", "", "The names of the parameters to delete.")
		ssm_deleteParametersCmd.MarkFlagRequired("names")
	})
	ssmCmd.AddCommand(ssm_deleteParametersCmd)
}
