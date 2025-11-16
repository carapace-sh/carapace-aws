package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_describeStudioCmd = &cobra.Command{
	Use:   "describe-studio",
	Short: "Returns details for the specified Amazon EMR Studio including ID, Name, VPC, Studio access URL, and so on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_describeStudioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_describeStudioCmd).Standalone()

		emr_describeStudioCmd.Flags().String("studio-id", "", "The Amazon EMR Studio ID.")
		emr_describeStudioCmd.MarkFlagRequired("studio-id")
	})
	emrCmd.AddCommand(emr_describeStudioCmd)
}
