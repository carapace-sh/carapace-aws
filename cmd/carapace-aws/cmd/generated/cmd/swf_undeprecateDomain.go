package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_undeprecateDomainCmd = &cobra.Command{
	Use:   "undeprecate-domain",
	Short: "Undeprecates a previously deprecated domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_undeprecateDomainCmd).Standalone()

	swf_undeprecateDomainCmd.Flags().String("name", "", "The name of the domain of the deprecated workflow type.")
	swf_undeprecateDomainCmd.MarkFlagRequired("name")
	swfCmd.AddCommand(swf_undeprecateDomainCmd)
}
