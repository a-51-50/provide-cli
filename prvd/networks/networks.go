package networks

import (
	"os"

	"github.com/provideplatform/provide-cli/prvd/common"
	"github.com/spf13/cobra"
)

var network map[string]interface{}
var networks []interface{}

var NetworksCmd = &cobra.Command{
	Use:   "networks",
	Short: "Manage networks",
	Long:  `Manage and provision elastic distributed networks and other infrastructure`,
	Run: func(cmd *cobra.Command, args []string) {
		common.CmdExistsOrExit(cmd, args)

		generalPrompt(cmd, args, "")

		defer func() {
			if r := recover(); r != nil {
				os.Exit(1)
			}
		}()
	},
}

func init() {
	NetworksCmd.AddCommand(networksInitCmd)
	NetworksCmd.AddCommand(networksListCmd)
	NetworksCmd.AddCommand(networksDisableCmd)
	NetworksCmd.Flags().BoolVarP(&optional, "optional", "", false, "List all the optional flags")
	NetworksCmd.Flags().BoolVarP(&paginate, "paginate", "", false, "List pagination flags")
}
