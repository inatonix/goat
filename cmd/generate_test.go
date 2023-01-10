package cmd_test

// import (
// 	"goat/cmd"
// 	"testing"
// )

// type parameter struct {
// 	ls string
// }
// type response struct {
// }

// func SearchTest(t *testing.T) {
// 	cases := map[string]struct {
// 		in   parameter
// 		want response
// 	}{
// 		"全てのパラメーターが使用される": {
// 			in: parameter{ls: `
// 				Dockerfile
// 				README.md
// 				batch
// 				config.yaml
// 				controller
// 				dao
// 				firebase.json
// 				go.mod
// 				go.sum
// 				logger
// 				main.go
// 				middleware
// 				model
// 				reference
// 				usecase
// 				util
// 			`},
// 			want: response{},
// 		},
// 	}
// 	for name, tt := range cases {
// 		tt := tt
// 		t.Run(name, func(t *testing.T) {
// 			t.Parallel()
// 			output, err := cmd.SearchHandlers(tt.in.ls)
// 			if err != nil {
// 			}

// 		})
// 	}
// }
