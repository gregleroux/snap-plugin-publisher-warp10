package warp10

import (
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestWarp10Plugin(t *testing.T) {

	Convey("Create Warp10 publisher", t, func() {
		gp := &Warp10Publisher{}
		Convey("So publisher should not be nil", func() {
			So(gp, ShouldNotBeNil)
		})

		Convey("Publisher should be of type warp10Publisher", func() {
			So(gp, ShouldHaveSameTypeAs, &Warp10Publisher{})
		})

		configPolicy, err := gp.GetConfigPolicy()
		Convey("Should return a config policy", func() {

			Convey("configPolicy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)

				Convey("and retrieving config policy should not error", func() {
					So(err, ShouldBeNil)

					Convey("config policy should be a cpolicy.ConfigPolicy", func() {
						So(configPolicy, ShouldHaveSameTypeAs, plugin.ConfigPolicy{})
					})

					testConfig := make(plugin.Config)
					testConfig["warp_url"] = "localhost"
					testConfig["token"] = "testtoken"

					server, err := testConfig.GetString("warp_url")
					Convey("So testConfig should return the right server config", func() {
						So(err, ShouldBeNil)
						So(server, ShouldEqual, "localhost")
					})

					port, err := testConfig.GetInt("token")
					Convey("So testConfig should return the right port config", func() {
						So(err, ShouldBeNil)
						So(server, ShouldEqual, "testtoken")
					})
				})
			})
		})
	})
}
