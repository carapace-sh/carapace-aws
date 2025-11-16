package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_describeProjectCmd = &cobra.Command{
	Use:   "describe-project",
	Short: "Returns the definition of a specific DataBrew project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_describeProjectCmd).Standalone()

	databrew_describeProjectCmd.Flags().String("name", "", "The name of the project to be described.")
	databrew_describeProjectCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_describeProjectCmd)
}
