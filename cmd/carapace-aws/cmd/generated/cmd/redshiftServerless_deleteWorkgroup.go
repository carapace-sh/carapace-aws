package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_deleteWorkgroupCmd = &cobra.Command{
	Use:   "delete-workgroup",
	Short: "Deletes a workgroup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_deleteWorkgroupCmd).Standalone()

	redshiftServerless_deleteWorkgroupCmd.Flags().String("workgroup-name", "", "The name of the workgroup to be deleted.")
	redshiftServerless_deleteWorkgroupCmd.MarkFlagRequired("workgroup-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_deleteWorkgroupCmd)
}
