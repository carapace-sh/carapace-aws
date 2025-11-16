package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_updateProjectCmd = &cobra.Command{
	Use:   "update-project",
	Short: "Modifies the definition of an existing DataBrew project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_updateProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_updateProjectCmd).Standalone()

		databrew_updateProjectCmd.Flags().String("name", "", "The name of the project to be updated.")
		databrew_updateProjectCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to be assumed for this request.")
		databrew_updateProjectCmd.Flags().String("sample", "", "")
		databrew_updateProjectCmd.MarkFlagRequired("name")
		databrew_updateProjectCmd.MarkFlagRequired("role-arn")
	})
	databrewCmd.AddCommand(databrew_updateProjectCmd)
}
