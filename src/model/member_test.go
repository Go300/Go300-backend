package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrivial(t *testing.T) {
	Convey("Username", t, func() {
		member := Member{
			Username: "MuslimBeibytuly",
		}
		newMember, err := Create(member)
		So(len(newMember.Username), ShouldEqual, 15)
		So(len(newMember.Id), ShouldEqual, 24)
		So(err, ShouldEqual, nil)
	})
}

func TestEmptyUsername(t *testing.T) {
	Convey("Username", t, func() {
		member := Member{
			Username: "",
		}
		_, err := Create(member)
		So(err, ShouldNotEqual, nil)
	})
}
