package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_listDbSystemShapesCmd = &cobra.Command{
	Use:   "list-db-system-shapes",
	Short: "Returns information about the shapes that are available for an Exadata infrastructure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_listDbSystemShapesCmd).Standalone()

	odb_listDbSystemShapesCmd.Flags().String("availability-zone", "", "The logical name of the AZ, for example, us-east-1a.")
	odb_listDbSystemShapesCmd.Flags().String("availability-zone-id", "", "The physical ID of the AZ, for example, use1-az4.")
	odb_listDbSystemShapesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	odb_listDbSystemShapesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	odbCmd.AddCommand(odb_listDbSystemShapesCmd)
}
