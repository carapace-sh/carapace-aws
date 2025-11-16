package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listAccessControlConfigurationsCmd = &cobra.Command{
	Use:   "list-access-control-configurations",
	Short: "Lists one or more access control configurations for an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listAccessControlConfigurationsCmd).Standalone()

	kendra_listAccessControlConfigurationsCmd.Flags().String("index-id", "", "The identifier of the index for the access control configuration.")
	kendra_listAccessControlConfigurationsCmd.Flags().String("max-results", "", "The maximum number of access control configurations to return.")
	kendra_listAccessControlConfigurationsCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there's more data to retrieve), Amazon Kendra returns a pagination token in the response.")
	kendra_listAccessControlConfigurationsCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_listAccessControlConfigurationsCmd)
}
