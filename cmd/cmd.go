package cmd

import (
	_ "github.com/ondrejsika/counter/cmd"
	counter_root "github.com/ondrejsika/counter/cmd/root"
	_ "github.com/ondrejsika/hellogophercamp/cmd"
	hellogophercamp_root "github.com/ondrejsika/hellogophercamp/cmd/root"
	_ "github.com/ondrejsika/nela-cli/cmd"
	nela_root "github.com/ondrejsika/nela-cli/cmd/root"
	_ "github.com/ondrejsika/training-cli/cmd"
	training_cli_root "github.com/ondrejsika/training-cli/cmd/root"
	_ "github.com/sikalabs/dogsay/cmd"
	dogsay_root "github.com/sikalabs/dogsay/cmd/root"
	_ "github.com/sikalabs/gobble/cmd"
	gobble_root "github.com/sikalabs/gobble/cmd/root"
	_ "github.com/sikalabs/install-slu/cmd"
	install_slu_root "github.com/sikalabs/install-slu/cmd/root"
	_ "github.com/sikalabs/mon/cmd"
	mon_root "github.com/sikalabs/mon/cmd/root"
	_ "github.com/sikalabs/redirect-server/cmd"
	redirect_server_root "github.com/sikalabs/redirect-server/cmd/root"
	_ "github.com/sikalabs/signpost/cmd"
	signpost_root "github.com/sikalabs/signpost/cmd/root"
	_ "github.com/sikalabs/slc/cmd"
	slc_root "github.com/sikalabs/slc/cmd/root"
	_ "github.com/sikalabs/slr/cmd"
	slr_root "github.com/sikalabs/slr/cmd/root"
	_ "github.com/sikalabs/slu/cmd"
	slu_root "github.com/sikalabs/slu/cmd/root"
	_ "github.com/sikalabs/tergum/cmd"
	tergum_root "github.com/sikalabs/tergum/cmd/root"
	"github.com/spf13/cobra"
)

var LargeCmd = &cobra.Command{
	Use:   "large",
	Short: "large: slu, gobble, signpost, ... in one binary",
}

func init() {
	LargeCmd.AddCommand(slu_root.RootCmd)
	LargeCmd.AddCommand(training_cli_root.Cmd)
	LargeCmd.AddCommand(signpost_root.Cmd)
	LargeCmd.AddCommand(tergum_root.Cmd)
	LargeCmd.AddCommand(install_slu_root.Cmd)
	LargeCmd.AddCommand(gobble_root.Cmd)
	LargeCmd.AddCommand(nela_root.Cmd)
	LargeCmd.AddCommand(mon_root.Cmd)
	LargeCmd.AddCommand(slc_root.RootCmd)
	LargeCmd.AddCommand(slr_root.Cmd)
	LargeCmd.AddCommand(redirect_server_root.Cmd)
	LargeCmd.AddCommand(dogsay_root.Cmd)
	LargeCmd.AddCommand(hellogophercamp_root.Cmd)
	LargeCmd.AddCommand(counter_root.Cmd)
}

func Execute() {
	LargeCmd.Execute()
}
