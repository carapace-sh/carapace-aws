package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getWorkgroupCmd = &cobra.Command{
	Use:   "get-workgroup",
	Short: "Returns information about a specific workgroup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getWorkgroupCmd).Standalone()

	redshiftServerless_getWorkgroupCmd.Flags().String("workgroup-name", "", "The name of the workgroup to return information for.")
	redshiftServerless_getWorkgroupCmd.MarkFlagRequired("workgroup-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_getWorkgroupCmd)
}
