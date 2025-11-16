package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listManagedJobTemplatesCmd = &cobra.Command{
	Use:   "list-managed-job-templates",
	Short: "Returns a list of managed job templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listManagedJobTemplatesCmd).Standalone()

	iot_listManagedJobTemplatesCmd.Flags().String("max-results", "", "Maximum number of entries that can be returned.")
	iot_listManagedJobTemplatesCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	iot_listManagedJobTemplatesCmd.Flags().String("template-name", "", "An optional parameter for template name.")
	iotCmd.AddCommand(iot_listManagedJobTemplatesCmd)
}
