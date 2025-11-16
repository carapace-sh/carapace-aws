package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getUsageProfileCmd = &cobra.Command{
	Use:   "get-usage-profile",
	Short: "Retrieves information about the specified Glue usage profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getUsageProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getUsageProfileCmd).Standalone()

		glue_getUsageProfileCmd.Flags().String("name", "", "The name of the usage profile to retrieve.")
		glue_getUsageProfileCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_getUsageProfileCmd)
}
