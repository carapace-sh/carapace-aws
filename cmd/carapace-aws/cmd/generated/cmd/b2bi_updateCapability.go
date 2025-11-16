package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_updateCapabilityCmd = &cobra.Command{
	Use:   "update-capability",
	Short: "Updates some of the parameters for a capability, based on the specified parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_updateCapabilityCmd).Standalone()

	b2bi_updateCapabilityCmd.Flags().String("capability-id", "", "Specifies a system-assigned unique identifier for the capability.")
	b2bi_updateCapabilityCmd.Flags().String("configuration", "", "Specifies a structure that contains the details for a capability.")
	b2bi_updateCapabilityCmd.Flags().String("instructions-documents", "", "Specifies one or more locations in Amazon S3, each specifying an EDI document that can be used with this capability.")
	b2bi_updateCapabilityCmd.Flags().String("name", "", "Specifies a new name for the capability, to replace the existing name.")
	b2bi_updateCapabilityCmd.MarkFlagRequired("capability-id")
	b2biCmd.AddCommand(b2bi_updateCapabilityCmd)
}
