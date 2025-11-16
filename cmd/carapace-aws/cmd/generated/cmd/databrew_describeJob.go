package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_describeJobCmd = &cobra.Command{
	Use:   "describe-job",
	Short: "Returns the definition of a specific DataBrew job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_describeJobCmd).Standalone()

	databrew_describeJobCmd.Flags().String("name", "", "The name of the job to be described.")
	databrew_describeJobCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_describeJobCmd)
}
