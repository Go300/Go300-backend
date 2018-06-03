package model

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)


func TestTrivial(t *testing.T) {
	Convey("Username", t, func() {
		member := Member {
			Name: "test name 1",
			Token: "",
		}
		new_member, err := Create(member)
		So(len(new_member.Token), ShouldEqual, 24)
		So(err, ShouldEqual, nil)
	})
}

func TestEmptyUsername(t *testing.T) {
	Convey("Username", t, func() {
		member := Member {
			Name: "",
			Token: "",
		}
		_, err := Create(member)
		So(err, ShouldNotEqual, nil)
	})
}