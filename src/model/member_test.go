package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTrivial(t *testing.T) {
	Convey("Username", t, func() {
		member := Member{
			Name:  "test name 1",
			Token: "",
		}
		newMember, err := Create(member)
		So(len(newMember.Token), ShouldEqual, 24)
		So(err, ShouldEqual, nil)
	})
}

func TestEmptyUsername(t *testing.T) {
	Convey("Username", t, func() {
		member := Member{
			Name:  "",
			Token: "",
		}
		_, err := Create(member)
		So(err, ShouldNotEqual, nil)
	})
}
