package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateUsageProfileCmd = &cobra.Command{
	Use:   "update-usage-profile",
	Short: "Update an Glue usage profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateUsageProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_updateUsageProfileCmd).Standalone()

		glue_updateUsageProfileCmd.Flags().String("configuration", "", "A `ProfileConfiguration` object specifying the job and session values for the profile.")
		glue_updateUsageProfileCmd.Flags().String("description", "", "A description of the usage profile.")
		glue_updateUsageProfileCmd.Flags().String("name", "", "The name of the usage profile.")
		glue_updateUsageProfileCmd.MarkFlagRequired("configuration")
		glue_updateUsageProfileCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_updateUsageProfileCmd)
}
